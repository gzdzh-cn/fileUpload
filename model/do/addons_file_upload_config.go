// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsFileUploadConfig is the golang structure of table addons_fileUpload_config for DAO operations like Where/Data.
type AddonsFileUploadConfig struct {
	g.Meta          `orm:"table:addons_fileUpload_config, do:true"`
	Id              interface{} //
	CreateTime      *gtime.Time // 创建时间
	UpdateTime      *gtime.Time // 更新时间
	DeletedAt       *gtime.Time //
	Name            interface{} // 站点名称
	Image           interface{} // 图片
	Link            interface{} // 跳转
	TypeId          interface{} // 类别
	Remark          interface{} // 备注
	Status          interface{} // 状态
	OrderNum        interface{} // 排序
	ItemId          interface{} // 站点编号
	SiteDomain      interface{} // 站点域名
	SiteIp          interface{} // IP白名单
	FtpHost         interface{} // ftp地址
	FtpPort         interface{} // ftp端口
	FtpUser         interface{} // ftp账号
	FtpPassword     interface{} // ftp密码
	LocalRootPath   interface{} // 本地根路径
	RemoteRoot      interface{} // 线上根路径
	LocalPathList   interface{} // 上传文件夹
	IgnoreList      interface{} // 忽略文件
	ProcessStatus   interface{} // 进度状态
	UploadTime      *gtime.Time //
	Percent         interface{} // 进度值
	UploadStartTime *gtime.Time //
	UploadEndTime   *gtime.Time //
	Error           interface{} // 异常信息
}
