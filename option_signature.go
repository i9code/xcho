package xcho

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"

	"github.com/i9code/xutil"
)

var (
	_ option = (*optionSignature)(nil)
)

type optionSignature struct {
	//  确定是不是要走中间件
	skipper middleware.Skipper
	//  签名算法
	algorithm xutil.Algorithm
	//  获得签名参数
	source keySource
}

// Signature Http签名
func Signature(algorithm xutil.Algorithm, source keySource) *optionSignature {
	return SignatureWithConfig(middleware.DefaultSkipper, algorithm, source)
}

// SignatureWithConfig Http签名
func SignatureWithConfig(skipper middleware.Skipper, algorithm xutil.Algorithm, source keySource) *optionSignature {
	// 检查算法配置是否正确
	if _, ok := xutil.AlgorithmMap[algorithm]; !ok {
		panic(fmt.Errorf("不支持的算法：%s", algorithm))
	}

	return &optionSignature{
		skipper:   skipper,
		algorithm: algorithm,
		source:    source,
	}
}

func (j *optionSignature) apply(options *options) {
	options.signature.skipper = j.skipper
	options.signature.algorithm = j.algorithm
	options.signature.source = j.source
	options.signatureEnable = true
}
