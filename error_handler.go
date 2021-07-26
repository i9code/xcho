package xcho

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/i9code/xcho/core"
	"github.com/i9code/xcho/lib/valid"
)

type errorHandler func(err error, ctx echo.Context)

func errorHandlerFunc(err error, ctx echo.Context) {
	rsp := core.CodeMessage{}

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
		lang := ctx.Request().Header.Get(core.HeaderAcceptLanguage)
		rsp.ErrorCode = 9901
		rsp.Message = "数据验证错误"
		rsp.Data = valid.I18n(lang, e)
	case *core.CodeMessage:
		rsp.ErrorCode = e.ToErrorCode()
		rsp.Message = e.ToMessage()
		rsp.Data = e.ToData()
	default:
		rsp.ErrorCode = 9903
		rsp.Message = "服务器内部错误"
		rsp.Data = err.Error()
	}

	_ = ctx.JSON(statusCode, rsp)
}
