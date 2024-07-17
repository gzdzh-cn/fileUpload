package admin

import (
	"context"
	v1 "dzhgo/addons/fileUpload/api/app_v1"
	"dzhgo/addons/fileUpload/dao"
	"dzhgo/addons/fileUpload/defineType"
	logic "dzhgo/addons/fileUpload/logic/sys"
	"dzhgo/addons/fileUpload/service"
	"fmt"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gzdzh/dzhcore"
)

type FileUploadConfigController struct {
	*dzhcore.Controller
}

func init() {

	var fileUploadConfigController = &FileUploadConfigController{
		&dzhcore.Controller{
			Prefix:  "/admin/fileUpload/config",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsFileUploadConfigService(),
		},
	}
	//注册路由
	dzhcore.RegisterController(fileUploadConfigController)

}

// 开始上传
func (c *FileUploadConfigController) StartUpload(ctx context.Context, req *v1.StartUploadReq) (res *dzhcore.BaseRes, err error) {

	//任务存在返回任务id
	id, _ := service.FileUploadService().GetProcessStatusById(ctx, req.ItemId)
	if id != "stop" {
		res = dzhcore.Ok(id)
		return
	}

	config := service.FileUploadService().GetConfig(ctx, req.ItemId)

	//检查目录是否存在
	if !gfile.Exists(config.LocalRootPath) {
		err = gerror.New("目录不存在")
		return
	}

	ftpConn, err := service.FileUploadService().ConnectToFtp(ctx, config, true)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}

	task := &defineType.Task{
		Id:        req.ItemId,
		Status:    "start",
		Percent:   0,
		Config:    config,
		FtpCon:    ftpConn,
		TaskQueue: gqueue.New(),
		SendData:  make(map[string]*defineType.SendData),
	}

	//添加任务
	service.TaskManager().AddTask(task)

	//执行上传
	service.FileUploadService().Upload(ctx, req.ItemId)

	data := fmt.Sprintf("项目:%v,本地执行路径:%v,远程上传路径:%v,启动成功", req.ItemId, config.LocalRootPath, config.RemoteRoot)
	res = dzhcore.Ok(data)
	return
}

// 停止上传
func (c *FileUploadConfigController) StopUploadById(ctx context.Context, req *v1.StopUploadByIdReq) (res *dzhcore.BaseRes, err error) {

	//任务不存在返回
	id, _ := service.FileUploadService().GetProcessStatusById(ctx, req.ItemId)
	if id == "stop" {
		return
	}

	data, err := service.FileUploadService().StopUploadById(ctx, req.ItemId)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 批量开始
func (c *FileUploadConfigController) MultiUpload(ctx context.Context, req *v1.MultiStartReq) (res *dzhcore.BaseRes, err error) {

	itemIds := g.SliceStr{}
	var errorData []*defineType.ErrorData

	for _, itemId := range req.ItemIds {

		//任务不存在 启动任务
		id, _ := service.FileUploadService().GetProcessStatusById(ctx, itemId)
		if id == "stop" {
			config := service.FileUploadService().GetConfig(ctx, itemId)

			//检查目录是否存在
			if !gfile.Exists(config.LocalRootPath) {
				err = gerror.New("目录不存在")
				g.Log().Errorf(ctx, "项目：%v,错误：%s", itemId, err)
				errorData = append(errorData, &defineType.ErrorData{
					ItemId: itemId,
					Error:  err.Error(),
				})
				continue
			}

			ftpConn, err := service.FileUploadService().ConnectToFtp(ctx, config, true)
			if err != nil {
				g.Log().Errorf(ctx, "项目：%v,错误：%s", itemId, err)
				errorData = append(errorData, &defineType.ErrorData{
					ItemId: itemId,
					Error:  err.Error(),
				})
				continue
			}

			task := &defineType.Task{
				Id:        itemId,
				Status:    "start",
				Percent:   0,
				Config:    config,
				FtpCon:    ftpConn,
				TaskQueue: gqueue.New(),
				SendData:  make(map[string]*defineType.SendData),
			}

			//添加任务
			service.TaskManager().AddTask(task)

			//执行上传
			service.FileUploadService().Upload(ctx, itemId)

			itemIds = append(itemIds, itemId)
		}
	}

	if len(errorData) > 0 {

		for _, item := range errorData {
			_, err = dao.AddonsFileUploadConfig.Ctx(ctx).Data(item).Where("itemId", item.ItemId).Update()
			if err != nil {
				g.Log().Errorf(ctx, "更新数据库，错误：%s", err)
			}
		}
	}

	res = dzhcore.Ok(errorData)
	return
}

// 批量停止
func (c *FileUploadConfigController) MultiStop(ctx context.Context, req *v1.MultiStopReq) (res *dzhcore.BaseRes, err error) {

	ids := g.SliceStr{}
	var errorData []*defineType.ErrorData
	for _, itemId := range req.ItemIds {

		//进度不存在就跳过
		id, _ := service.FileUploadService().GetProcessStatusById(ctx, itemId)
		if id == "stop" {
			continue
		}
		_, err = service.FileUploadService().StopUploadById(ctx, itemId)
		if err != nil {
			g.Log().Errorf(ctx, "项目：%v,错误：%s", itemId, err)
			errorData = append(errorData, &defineType.ErrorData{
				ItemId: itemId,
				Error:  err.Error(),
			})
			continue
		}
		ids = append(ids, itemId)
	}

	data := fmt.Sprintf("项目:%v,停止成功", ids)
	res = dzhcore.Ok(data)
	return
}

// 获取进度状态
func (c *FileUploadConfigController) GetProcessStatusById(ctx context.Context, req *v1.GetProcessStatusByIdReq) (res *dzhcore.BaseRes, err error) {

	data, err := service.FileUploadService().GetProcessStatusById(ctx, req.ItemId)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
