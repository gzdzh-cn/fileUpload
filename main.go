package fileUpload

import (
	"context"
	"dzhgo/addons/fileUpload/service"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhcore"
	"net/http"
	"time"

	baseModel "dzhgo/internal/model"

	_ "dzhgo/addons/fileUpload/controller"
	_ "dzhgo/addons/fileUpload/funcs"
	_ "dzhgo/addons/fileUpload/middleware"
	_ "dzhgo/addons/fileUpload/packed"
)

var ctx = gctx.GetInitCtx()

// 流返回
func steamRes(ctx context.Context, r *ghttp.Request, itemId string) {
	//  流式回应
	rw := r.Response.RawWriter()
	flusher, ok := rw.(http.Flusher)
	if !ok {
		g.Log().Error(ctx, "rw.(http.Flusher) error")
		r.Response.WriteStatusExit(500)
		return
	}
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	// 通过循环每隔一段时间发送一次数据
	for {

		sendData, err := service.FileUploadService().GetProcessById(ctx, itemId)

		if err != nil {
			return
		}
		_, err = fmt.Fprintf(rw, "%s\n", gconv.String(dzhcore.Ok(sendData)))
		if err != nil {
			return
		}
		//g.Log().Warningf(ctx, "流循环:%v", gconv.String(task.SendData))

		if sendData.Status == false {
			break
		}

		// 刷新缓冲区，将数据发送给客户端
		flusher.Flush()
		time.Sleep(time.Millisecond * 500)

	}

	// 发送完成信号
	//_, _ = fmt.Fprintf(rw, "data: Task Completed\n\n")
	// 刷新缓冲区，确保最后的数据发送给客户端
	flusher.Flush()
}

func NewInit() {

	g.Log().Debug(ctx, "addon fileUpload init start ...")
	g.Log().Debugf(ctx, "fileUpload version:%v", Version)
	dzhcore.FillInitData(ctx, "fileUpload", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "addon fileUpload init finished ...")

	s := g.Server()
	s.BindHandler("/admin/fileUpload/steam/getProcessById", func(r *ghttp.Request) {

		ctx = gctx.New()
		itemId := r.Get("itemId").String()
		g.Log().Debug(ctx, "接收steam:", itemId)

		steamRes(ctx, r, itemId)

	})

	s.BindHandler("/app/fileUpload/steam/getProcessById", func(r *ghttp.Request) {

		ctx = gctx.New()
		itemId := r.Get("itemId").String()
		g.Log().Debug(ctx, "接收steam:", itemId)

		steamRes(ctx, r, itemId)

	})

}
