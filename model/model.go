package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.GetInitCtx()

func NewInit() {
	g.Log().Debug(ctx, "fileUpload init model")
}
