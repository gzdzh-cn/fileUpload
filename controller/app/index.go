package app

import (
	"context"
	"github.com/gzdzh/dzhcore"

	v1 "dzhgo/addons/fileUpload/api/app_v1"
	"dzhgo/addons/fileUpload/defineType"
	"dzhgo/addons/fileUpload/service"
	"fmt"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

type IndexController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var indexController = &IndexController{
		&dzhcore.ControllerSimple{
			Prefix: "/app/fileUpload",
		},
	}
	// 注册路由
	dzhcore.RegisterControllerSimple(indexController)
}

// 开始上传
func (c *IndexController) StartUpload(ctx context.Context, req *v1.StartUploadReq) (res *dzhcore.BaseRes, err error) {

	//进度存在返回任务id
	result, err := service.FileUploadService().GetProcessStatusById(ctx, req.ItemId)
	if err != nil {
		res = dzhcore.Ok(result, err.Error())
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
func (s *IndexController) StopUploadById(ctx context.Context, req *v1.StopUploadByIdReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.FileUploadService().StopUploadById(ctx, req.ItemId)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 全部停止上传
func (s *IndexController) StopUploadByAll(ctx context.Context, req *v1.StopUploadByAllReq) (res *dzhcore.BaseRes, err error) {

	return
}

// 获取进度
func (c *IndexController) GetProcessStatusById(ctx context.Context, req *v1.GetProcessStatusByIdReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.FileUploadService().GetProcessStatusById(ctx, req.ItemId)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
