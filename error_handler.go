package xcho

import (
	`github.com/i9code/xutil/base`
	`net/http`

	`github.com/go-playground/validator/v10`
	`github.com/i9code/xutil/valid`
	`github.com/labstack/echo/v4`
)

type errorHandler func(err error, ctx echo.Context)

func errorHandlerFunc(err error, ctx echo.Context) {
	rsp := response{}

	statusCode := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		statusCode = e.Code
		rsp.ErrorCode = 9902
		rsp.Message = "处理请求失败"
		if nil != e.Internal {
			rsp.Data = e.Internal.Error()
		}
	case validator.ValidationErrors:
		statusCode = http.StatusBadRequest
		lang := ctx.Request().Header.Get(HeaderAcceptLanguage)
		rsp.ErrorCode = 9901
		rsp.Message = "数据验证错误"
		rsp.Data = valid.I18n(lang, e)
	case *base.ErrorCode:
		rsp.ErrorCode = e.Code()
		rsp.Message = e.Msg()
		rsp.Data = e.ToData()
	default:
		rsp.ErrorCode = 9903
		rsp.Message = "服务器内部错误"
		rsp.Data = err.Error()
	}

	_ = ctx.JSON(statusCode, rsp)
}
