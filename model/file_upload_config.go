package model

import (
	"github.com/gzdzh/dzhcore"
	"time"
)

const TableNameFileUploadConfig = "addons_file_upload_config"

// FileUploadConfig mapped from table <member_open>
type FileUploadConfig struct {
	*dzhcore.Model
	Name     string  `gorm:"column:name;not null;comment:站点名称" json:"name"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Link     *string `gorm:"column:link;comment:跳转" json:"link"`
	TypeId   int64   `gorm:"column:typeId;comment:类别;index" json:"typeId"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int;default:1" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`

	ItemId          string    `gorm:"column:itemId;not null;uniqueIndex;primaryKey;varchar(255);comment:站点编号;" json:"itemId"`
	SiteDomain      string    `gorm:"column:siteDomain;comment:站点域名;" json:"siteDomain"`
	SiteIp          string    `gorm:"column:siteIp;comment:IP白名单;" json:"siteIp"`
	FtpHost         string    `gorm:"column:ftpHost;comment:ftp地址;" json:"ftpHost"`                              // ftp地址
	FtpPort         string    `gorm:"column:ftpPort;comment:ftp端口;" json:"ftpPort"`                              // ftp端口
	FtpUser         string    `gorm:"column:ftpUser;comment:ftp账号;" json:"ftpUser"`                              // ftp账号
	FtpPassword     string    `gorm:"column:ftpPassword;comment:ftp密码;" json:"ftpPassword"`                      // ftp密码
	LocalRootPath   string    `gorm:"column:localRootPath;comment:本地根路径;" json:"localRootPath"`                  // 本地根路径
	RemoteRoot      string    `gorm:"column:remoteRoot;comment:线上根路径;" json:"remoteRoot"`                        // 线上根路径
	LocalPathList   string    `gorm:"column:localPathList;comment:上传文件夹;" json:"localPathList"`                  // 要上传的文件夹或者文件
	IgnoreList      string    `gorm:"column:ignoreList;comment:忽略文件;" json:"ignoreList"`                         // 忽略文件
	ProcessStatus   string    `gorm:"column:processStatus;comment:进度状态;type:int;default:0" json:"processStatus"` //进度状态,0未执行 1执行中 2执行完成 -1停止 -2异常
	Percent         string    `gorm:"column:percent;comment:进度值;type:float;default:0" json:"percent"`            //进度值
	UploadStartTime time.Time `gorm:"column:uploadStartTime;index,comment:执行开始时间" json:"uploadStartTime"`        // 执行开始时间
	UploadEndTime   time.Time `gorm:"column:uploadEndTime;index,comment:执行结束时间" json:"uploadEndTime"`            // 执行结束时间
	Error           *string   `gorm:"column:error;comment:异常信息;" json:"error"`                                   //异常信息
}

// TableName FileUploadConfig's table name
func (*FileUploadConfig) TableName() string {
	return TableNameFileUploadConfig
}

// GroupName FileUploadConfig's table group
func (*FileUploadConfig) GroupName() string {
	return "default"
}

// NewFileUploadConfig create a new FileUploadConfig
func NewFileUploadConfig() *FileUploadConfig {
	return &FileUploadConfig{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	err := dzhcore.CreateTable(&FileUploadConfig{})
	if err != nil {
		return
	}
}
