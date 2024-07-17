package defineType

import (
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	ftp "github.com/gzdzh-cn/dzh-ftp"
	"sync"
)

// 配置参数
type Config struct {
	FtpHost       string     `json:"ftpHost"`       // ftp地址
	FtpPort       string     `json:"ftpPort"`       // ftp端口
	FtpUser       string     `json:"ftpUser"`       // ftp账号
	FtpPassword   string     `json:"ftpPassword"`   // ftp密码
	LocalRootPath string     `json:"localRootPath"` // 上传文件的根目录路径
	RemoteRoot    string     `json:"remoteRoot"`    // 根目录
	LocalPathList g.SliceStr `json:"localPathList"` // 要上传的文件夹或者文件
	IgnoreList    g.SliceStr `json:"ignoreList"`    // 忽略文件
	ItemId        string     `json:"itemId"`        // 站点编号
	SiteDomain    string     `json:"siteDomain"`    // 域名
	SiteId        string     `json:"siteId"`        // 白名单IP
}

// 发送的数据
type SendData struct {
	sync.Mutex
	ItemId     string  `json:"itemId"`
	Total      int     `json:"total"`
	HasSendNum int     `json:"hasSendNum"`
	Percent    float64 `json:"percent"`
	SendFile   string  `json:"sendFile"`
	Status     bool    `json:"status"`
}

// 任务
type Task struct {
	sync.Mutex
	Id        string               `json:"id"`
	Name      string               `json:"name"`      // 任务名称
	Config    *Config              `json:"config"`    // 上传服务器的配置
	SendData  map[string]*SendData `json:"sendData"`  // 要发送的数据
	FtpCon    *ftp.ServerConn      `json:"ftpCon"`    // ftp连接
	TaskQueue *gqueue.Queue        `json:"taskQueue"` // 队列
	Total     int                  `json:"total"`     //总文件数
	Percent   float64              `json:"percent"`   // 任务进度
	Status    string               `json:"status"`    // 任务状态
}

// 记录错误
type ErrorData struct {
	ItemId string `json:"itemId"`
	Error  string `json:"error"`
}
