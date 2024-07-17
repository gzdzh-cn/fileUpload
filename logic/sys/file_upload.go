package sys

import (
	"context"
	"dzhgo/addons/fileUpload/dao"
	"dzhgo/addons/fileUpload/defineType"
	"dzhgo/addons/fileUpload/service"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzh-ftp"

	"os"
	"path/filepath"
	"time"
)

type sFileUploadService struct{}

var (
	stopSignals = make(map[string]chan string)
)

func init() {
	service.RegisterFileUploadService(&sFileUploadService{})
}

// 实例化配置
func (s *sFileUploadService) GetConfig(ctx context.Context, itemId string) *defineType.Config {

	siteData, _ := dao.AddonsFileUploadConfig.Ctx(ctx).Where(g.Map{"itemId": itemId}).One()

	config := &defineType.Config{}
	gconv.Struct(siteData, config)

	localPathList := gstr.Split(gstr.TrimAll(gconv.String(siteData["localPathList"])), ",")
	ignoreList := gstr.Split(gstr.TrimAll(gconv.String(siteData["ignoreList"])), ",")

	config.LocalPathList = localPathList
	config.IgnoreList = ignoreList

	g.Log().Debugf(ctx, "config：%s", gconv.String(config))
	return config
}

// 连接ftp
func (s *sFileUploadService) ConnectToFtp(ctx context.Context, config *defineType.Config, mode bool) (*ftp.ServerConn, error) {

	address := fmt.Sprintf("%s:%s", config.FtpHost, config.FtpPort)
	if mode {
		g.Log().Warning(ctx, "主动模式")
	} else {
		g.Log().Warning(ctx, "被动模式")
	}
	conn, err := ftp.Dial(address, ftp.DialWithTimeout(5*time.Second), ftp.DialWithDisabledEPSV(mode))
	if err != nil {
		return nil, err
	}

	err = conn.Login(config.FtpUser, config.FtpPassword)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// 执行上传
func (s *sFileUploadService) Upload(ctx context.Context, itemId string) {

	task := service.TaskManager().GetTask(itemId)
	stopSignals[itemId] = make(chan string, 1)

	task.SendData[itemId] = &defineType.SendData{}

	g.Log().Warningf(ctx, "任务：%v，开始", itemId)
	// 使用一个新的背景上下文
	newCtx := gctx.New()
	go func() {

		//记录开始
		updateData := g.Map{
			"processStatus":   1,
			"uploadStartTime": gtime.Now(),
			"error":           nil,
		}

		_, err := dao.AddonsFileUploadConfig.Ctx(newCtx).Data(updateData).Where(g.Map{"itemId": itemId}).Update()
		if err != nil {
			g.Log().Error(newCtx, err.Error())
			service.TaskManager().DelTask(itemId)
			return
		}

		//计算全部文件
		list, err := getFileList(ctx, task)
		if err != nil {
			g.Log().Error(ctx, err)
			service.TaskManager().DelTask(itemId)
			updateData := g.Map{
				"processStatus": -2,
				"error":         err.Error(),
			}
			_, err = dao.AddonsFileUploadConfig.Ctx(newCtx).Data(updateData).Where(g.Map{"itemId": itemId}).Update()
			if err != nil {
				g.Log().Error(newCtx, err.Error())
				return
			}
		}
		task.Total = len(list)

		//队列输入
		pubQueue(newCtx, task.TaskQueue, list)
	}()

	go func() {

		//队列输出
		err := popQueue(newCtx, task.TaskQueue, task)
		if err != nil {
			g.Log().Error(ctx, err)
			service.TaskManager().DelTask(itemId)
			updateData := g.Map{
				"processStatus": -2,
				"error":         err.Error(),
			}
			_, err = dao.AddonsFileUploadConfig.Ctx(newCtx).Data(updateData).Where(g.Map{"itemId": itemId}).Update()
			if err != nil {
				g.Log().Error(newCtx, err.Error())
			}
		}
	}()

}

// 队列输入
func pubQueue(ctx context.Context, taskQueue *gqueue.Queue, fileList g.SliceStr) {

	for _, file := range fileList {

		//写入队列
		taskQueue.Push(file)
		g.Log().Debugf(ctx, "队列输入：%v", file)
		//time.Sleep(time.Second)
	}

}

// 队列输出
func popQueue(ctx context.Context, taskQueue *gqueue.Queue, task *defineType.Task) error {

	var num = 0
	for {

		select {
		//收到信号，停止循环
		case <-stopSignals[task.Id]:

			//记录停止
			updateData := g.Map{
				"processStatus": -1,
				"percent":       task.SendData[task.Id].Percent,
				"uploadEndTime": gtime.Now(),
			}
			_, err := dao.AddonsFileUploadConfig.Ctx(ctx).Data(updateData).Where(g.Map{"itemId": task.Id}).Update()
			if err != nil {
				return err
			}

			//task.SendData[task.Id].Lock()
			//task.SendData[task.Id].Status = false
			//task.SendData[task.Id].Unlock()
			service.TaskManager().DelTask(task.Id)

			task.Lock()
			task.Status = "end"
			task.Unlock()
			g.Log().Warningf(ctx, "退出队列：%s", task.Id)

			return nil

		default:

			//队列输出
			pop := taskQueue.Pop()
			if pop == nil {
				glog.Debug(ctx, "队列为空，继续等待")
				time.Sleep(time.Second) // 避免忙等待
				continue
			}
			g.Log().Debugf(ctx, "队列输出：%v", gconv.String(pop))

			file := gconv.String(pop)
			//上传
			err := ftpUpload(ctx, task, file)
			if err != nil {
				return err
			}

			num++
			//计算百分比
			percent := computerProgress(ctx, num, task.Total)

			task.SendData[task.Id].Lock()
			task.SendData[task.Id].ItemId = task.Id
			task.SendData[task.Id].Total = task.Total
			task.SendData[task.Id].SendFile = file
			task.SendData[task.Id].Percent = percent
			task.SendData[task.Id].Status = true
			task.SendData[task.Id].HasSendNum = num
			task.SendData[task.Id].Unlock()

			task.Lock()
			task.Percent = percent
			task.Unlock()

			if task.SendData[task.Id].Percent == 100 {

				task.SendData[task.Id].Lock()
				task.Status = "end"
				task.SendData[task.Id].Unlock()

				task.SendData[task.Id].Lock()
				task.SendData[task.Id].Status = false
				task.SendData[task.Id].Unlock()

				updateData := g.Map{
					"processStatus": 2,
					"percent":       task.SendData[task.Id].Percent,
					"uploadEndTime": gtime.Now(),
				}
				_, err = dao.AddonsFileUploadConfig.Ctx(ctx).Data(updateData).Where(g.Map{"itemId": task.Id}).Update()
				if err != nil {
					return err
				}
			}
			//glog.Debugf(ctx, "上传中：num：%v", gconv.String(num))
			g.Log().Debugf(ctx, "上传中：%v", gconv.String(task.SendData[task.Id]))
		}

		//time.Sleep(time.Millisecond * 200)
	}
}

// 上传文件
func ftpUpload(ctx context.Context, task *defineType.Task, filepath string) error {

	config := task.Config
	conn := task.FtpCon
	SendFile := config.LocalRootPath + filepath
	remoteFile := config.RemoteRoot + filepath

	//如果文件夹就创建
	if gfile.Ext(remoteFile) == "" {
		err := createRemoteDir(ctx, conn, remoteFile)
		if err != nil {
			g.Log().Warning(ctx, err.Error())
		}
	} else {
		//如果是文件就上传
		err := uploadFile(ctx, conn, SendFile, remoteFile)
		if err != nil {
			return err
		}
	}

	return nil
}

// 全部文件列表（含目录）
func getFileList(ctx context.Context, task *defineType.Task) (list g.SliceStr, err error) {

	g.Log().Debug(ctx, "getFileList start")

	for _, localPath := range task.Config.LocalPathList {

		g.Log().Debugf(ctx, "localPath:%v", localPath)

		//检查是否存在
		if !gfile.Exists(task.Config.LocalRootPath + localPath) {
			err = gerror.Newf("%v不存在", localPath)
			return
		}
		// ignores some files
		ignores := garray.NewStrArrayFrom(task.Config.IgnoreList)

		if !ignores.Contains(gfile.Basename(localPath)) {
			list = append(list, localPath)
		}

		if gfile.Ext(localPath) == "" {
			//扫描目录，过滤自定义文件，得到需要上传的文件夹和文件集合
			fileList, err := gfile.ScanDirFunc(task.Config.LocalRootPath+localPath, "*", true, func(path string) string {
				if ignores.Contains(gfile.Basename(path)) {
					return ""
				}
				path_ := gstr.SubStrFrom(path, localPath)
				return path_
			})
			if err != nil {
				return nil, err
			}
			list = append(list, fileList...)
		}
	}
	g.Log().Debugf(ctx, "getFileList end : %v ", len(list))

	return
}

// 计算百分比
func computerProgress(ctx context.Context, sendNum int, total int) float64 {

	num := gconv.Float64((gconv.Float64(sendNum) / gconv.Float64(total)) * 100)
	return gconv.Float64(fmt.Sprintf("%.2f", num))
}

// 指定任务id停止
func (s *sFileUploadService) StopUploadById(ctx context.Context, itemId string) (data interface{}, err error) {

	// 发送停止信号
	stopById(ctx, itemId)
	return fmt.Sprintf("停止队列：%v", itemId), nil
}

// 全部任务停止
func (s *sFileUploadService) StopUploadByAll(ctx context.Context) (data interface{}, err error) {

	stopAll(ctx)
	g.Log().Debug(ctx, "全部队列停止")
	return fmt.Sprint("全部队列停止"), nil
}

// 根据itemId获取进度（流返回使用）
func (s *sFileUploadService) GetProcessById(ctx context.Context, itemId string) (data *defineType.SendData, err error) {

	task := service.TaskManager().GetTask(itemId)
	//glog.Debugf(ctx, "流返回：%v", gconv.String(task.SendData))
	return task.SendData[task.Id], nil
}

// 根据itemId获取进度（不是流返回）
func (s *sFileUploadService) GetProcessStatusById(ctx context.Context, itemId string) (data string, err error) {

	task := service.TaskManager().GetTask(itemId)
	if task == nil || task.Status == "end" {

		str := fmt.Sprintf("项目：%v,没启动", itemId)
		return str, nil
	} else {
		err = gerror.Newf("项目：%v,启动了", itemId)
		return task.Id, err
	}
}

// 指定任务id停止
func stopById(ctx context.Context, itemId string) {

	// 发送停止信号
	stopSignals[itemId] <- itemId
	g.Log().Debugf(ctx, "停止队列：%v", itemId)
}

// 停止全部队列
func stopAll(ctx context.Context) {

}

// ftp创建文件夹
func createRemoteDir(ctx context.Context, conn *ftp.ServerConn, path string) error {

	paths := filepath.SplitList(path)
	cumulativePath := ""
	for _, part := range paths {
		cumulativePath = filepath.Join(cumulativePath, part)
		err := conn.MakeDir(cumulativePath)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}

// 上传
func uploadFile(ctx context.Context, conn *ftp.ServerConn, sendFile string, remotePath string) error {

	file, err := gfile.Open(sendFile)
	if err != nil {
		return err
	}
	defer file.Close()

	g.Log().Debugf(ctx, "上传本地文件SendFile：%v,到线上remotePath：%v", sendFile, remotePath)
	err = conn.Stor(remotePath, file)
	if err != nil {
		return err
	}
	return nil
}
