// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsFileUploadConfigDao is the data access object for table addons_fileUpload_config.
type AddonsFileUploadConfigDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns AddonsFileUploadConfigColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsFileUploadConfigColumns defines and stores column names for table addons_fileUpload_config.
type AddonsFileUploadConfigColumns struct {
	Id              string //
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
	DeletedAt       string //
	Name            string // 站点名称
	Image           string // 图片
	Link            string // 跳转
	TypeId          string // 类别
	Remark          string // 备注
	Status          string // 状态
	OrderNum        string // 排序
	ItemId          string // 站点编号
	SiteDomain      string // 站点域名
	SiteIp          string // IP白名单
	FtpHost         string // ftp地址
	FtpPort         string // ftp端口
	FtpUser         string // ftp账号
	FtpPassword     string // ftp密码
	LocalRootPath   string // 本地根路径
	RemoteRoot      string // 线上根路径
	LocalPathList   string // 上传文件夹
	IgnoreList      string // 忽略文件
	ProcessStatus   string // 进度状态
	UploadTime      string //
	Percent         string // 进度值
	UploadStartTime string //
	UploadEndTime   string //
	Error           string // 异常信息
}

// addonsFileUploadConfigColumns holds the columns for table addons_fileUpload_config.
var addonsFileUploadConfigColumns = AddonsFileUploadConfigColumns{
	Id:              "id",
	CreateTime:      "createTime",
	UpdateTime:      "updateTime",
	DeletedAt:       "deleted_at",
	Name:            "name",
	Image:           "image",
	Link:            "link",
	TypeId:          "typeId",
	Remark:          "remark",
	Status:          "status",
	OrderNum:        "orderNum",
	ItemId:          "itemId",
	SiteDomain:      "siteDomain",
	SiteIp:          "siteIp",
	FtpHost:         "ftpHost",
	FtpPort:         "ftpPort",
	FtpUser:         "ftpUser",
	FtpPassword:     "ftpPassword",
	LocalRootPath:   "localRootPath",
	RemoteRoot:      "remoteRoot",
	LocalPathList:   "localPathList",
	IgnoreList:      "ignoreList",
	ProcessStatus:   "processStatus",
	UploadTime:      "uploadTime",
	Percent:         "percent",
	UploadStartTime: "uploadStartTime",
	UploadEndTime:   "uploadEndTime",
	Error:           "error",
}

// NewAddonsFileUploadConfigDao creates and returns a new DAO object for table data access.
func NewAddonsFileUploadConfigDao() *AddonsFileUploadConfigDao {
	return &AddonsFileUploadConfigDao{
		group:   "default",
		table:   "addons_fileUpload_config",
		columns: addonsFileUploadConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsFileUploadConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsFileUploadConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsFileUploadConfigDao) Columns() AddonsFileUploadConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsFileUploadConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsFileUploadConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsFileUploadConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
