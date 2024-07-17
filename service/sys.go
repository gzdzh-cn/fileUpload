// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"dzhgo/addons/fileUpload/defineType"

	ftp "github.com/gzdzh-cn/dzh-ftp"
)

type (
	IConfigService     interface{}
	IFileUploadService interface {
		// 实例化配置
		GetConfig(ctx context.Context, itemId string) *defineType.Config
		// 连接ftp
		ConnectToFtp(ctx context.Context, config *defineType.Config, mode bool) (*ftp.ServerConn, error)
		// 执行上传
		Upload(ctx context.Context, itemId string)
		// 指定任务id停止
		StopUploadById(ctx context.Context, itemId string) (data interface{}, err error)
		// 全部任务停止
		StopUploadByAll(ctx context.Context) (data interface{}, err error)
		// 根据itemId获取进度（流返回使用）
		GetProcessById(ctx context.Context, itemId string) (data *defineType.SendData, err error)
		// 根据itemId获取进度（不是流返回）
		GetProcessStatusById(ctx context.Context, itemId string) (data string, err error)
	}
	ITaskManager interface {
		// 添加任务
		AddTask(task *defineType.Task)
		// 删除任务
		DelTask(id string)
		// 获取任务
		GetTask(id string) *defineType.Task
		// 更新任务
		UpdateTask(task *defineType.Task)
	}
)

var (
	localConfigService     IConfigService
	localFileUploadService IFileUploadService
	localTaskManager       ITaskManager
)

func ConfigService() IConfigService {
	if localConfigService == nil {
		panic("implement not found for interface IConfigService, forgot register?")
	}
	return localConfigService
}

func RegisterConfigService(i IConfigService) {
	localConfigService = i
}

func FileUploadService() IFileUploadService {
	if localFileUploadService == nil {
		panic("implement not found for interface IFileUploadService, forgot register?")
	}
	return localFileUploadService
}

func RegisterFileUploadService(i IFileUploadService) {
	localFileUploadService = i
}

func TaskManager() ITaskManager {
	if localTaskManager == nil {
		panic("implement not found for interface ITaskManager, forgot register?")
	}
	return localTaskManager
}

func RegisterTaskManager(i ITaskManager) {
	localTaskManager = i
}
