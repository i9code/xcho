package xcho

import (
	"github.com/i9code/xutil"
	"github.com/labstack/echo/v4/middleware"
)

type signatureConfig struct {
	//  确定是不是要走中间件
	skipper middleware.Skipper `validate:"required"`
	//  签名算法
	algorithm xutil.Algorithm `validate:"required"`
	//  获得签名参数
	source keySource `validate:"required"`
}
