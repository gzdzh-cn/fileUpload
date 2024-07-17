package app_v1

import (
	"dzhgo/addons/fileUpload/defineType"
	"github.com/gogf/gf/v2/frame/g"
)

// 进度状态
type GetProcessStatusByIdReq struct {
	g.Meta `path:"/getProcessStatusById" method:"POST"`
	ItemId string `json:"itemId" v:"required#请输入itemId"` // 项目id
}

// 打印队列
type GetQueueByIdReq struct {
	g.Meta `path:"/getQueueById" method:"POST"`
	ItemId string `json:"itemId"` // 项目id
}

// 启动上传
type StartUploadReq struct {
	g.Meta `path:"/startUpload" method:"POST"`
	ItemId string             `json:"itemId"` // 项目id
	Config *defineType.Config `json:"config"`
}

// 指定id停止上传
type StopUploadByIdReq struct {
	g.Meta `path:"/stopUploadById" method:"POST"`
	ItemId string `json:"itemId"` // 项目id
}

// 全部停止上传
type StopUploadByAllReq struct {
	g.Meta `path:"/stopUploadByAll" method:"POST"`
}

// 批量开始
type MultiStartReq struct {
	g.Meta  `path:"/multiStart" method:"POST"`
	ItemIds g.SliceStr `json:"itemIds"` // 项目id 数组
}

// 批量停止
type MultiStopReq struct {
	g.Meta  `path:"/multiStop" method:"POST"`
	ItemIds g.SliceStr `json:"itemIds"` // 项目id 数组
}
