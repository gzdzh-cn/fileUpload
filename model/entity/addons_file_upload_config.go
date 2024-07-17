// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsFileUploadConfig is the golang structure for table addons_file_upload_config.
type AddonsFileUploadConfig struct {
	Id              string      `json:"id"              orm:"id"              ` //
	CreateTime      *gtime.Time `json:"createTime"      orm:"createTime"      ` // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"      orm:"updateTime"      ` // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"      ` //
	Name            string      `json:"name"            orm:"name"            ` // 站点名称
	Image           string      `json:"image"           orm:"image"           ` // 图片
	Link            string      `json:"link"            orm:"link"            ` // 跳转
	TypeId          int64       `json:"typeId"          orm:"typeId"          ` // 类别
	Remark          string      `json:"remark"          orm:"remark"          ` // 备注
	Status          int         `json:"status"          orm:"status"          ` // 状态
	OrderNum        int         `json:"orderNum"        orm:"orderNum"        ` // 排序
	ItemId          string      `json:"itemId"          orm:"itemId"          ` // 站点编号
	SiteDomain      string      `json:"siteDomain"      orm:"siteDomain"      ` // 站点域名
	SiteIp          string      `json:"siteIp"          orm:"siteIp"          ` // IP白名单
	FtpHost         string      `json:"ftpHost"         orm:"ftpHost"         ` // ftp地址
	FtpPort         string      `json:"ftpPort"         orm:"ftpPort"         ` // ftp端口
	FtpUser         string      `json:"ftpUser"         orm:"ftpUser"         ` // ftp账号
	FtpPassword     string      `json:"ftpPassword"     orm:"ftpPassword"     ` // ftp密码
	LocalRootPath   string      `json:"localRootPath"   orm:"localRootPath"   ` // 本地根路径
	RemoteRoot      string      `json:"remoteRoot"      orm:"remoteRoot"      ` // 线上根路径
	LocalPathList   string      `json:"localPathList"   orm:"localPathList"   ` // 上传文件夹
	IgnoreList      string      `json:"ignoreList"      orm:"ignoreList"      ` // 忽略文件
	ProcessStatus   int         `json:"processStatus"   orm:"processStatus"   ` // 进度状态
	Percent         float64     `json:"percent"         orm:"percent"         ` // 进度值
	UploadStartTime *gtime.Time `json:"uploadStartTime" orm:"uploadStartTime" ` //
	UploadEndTime   *gtime.Time `json:"uploadEndTime"   orm:"uploadEndTime"   ` //
	Error           string      `json:"error"           orm:"error"           ` // 异常信息
}
