package defineType

import (
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	ftp "github.com/gzdzh/dzh-ftp"
	"sync"
)

// 配置参数
type Config struct {
	FtpHost       string     // ftp地址
	FtpPort       string     // ftp端口
	FtpUser       string     // ftp账号
	FtpPassword   string     // ftp密码
	LocalRootPath string     // 上传文件的根目录路径
	RemoteRoot    string     // 根目录
	LocalPathList g.SliceStr // 要上传的文件夹或者文件
	IgnoreList    g.SliceStr // 忽略文件
	ItemId        string     // 站点编号
	SiteDomain    string     // 域名
	SiteId        string     // 白名单IP
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
	Id        string
	Name      string               // 任务名称
	Config    *Config              // 上传服务器的配置
	SendData  map[string]*SendData // 要发送的数据
	FtpCon    *ftp.ServerConn      // ftp连接
	TaskQueue *gqueue.Queue        // 队列
	Total     int                  //总文件数
	Percent   float64              // 任务进度
	Status    string               // 任务状态
}

// 记录错误
type ErrorData struct {
	ItemId string `json:"itemId"`
	Error  string `json:"error"`
}
