package base

import "github.com/i9code/xcho/base/xjson"

type (
	// 接口，符合条件的错误统一处理
	Error interface {
		// 返回错误码
		Code() int
		// 返回错误消息
		Msg() string
		// ToData 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		ToData() interface{}
	}

	// 带错误编号和信息的
	ErrorCode struct {
		// ErrorCode 错误码
		ErrorCode int `json:"errorCode"`
		// Message 消息
		Message string `json:"message"`
		// Data 数据
		Data interface{} `json:"data"`
	}
)

//  创建错误
func NewErrorCode(errorCode int, message string, data interface{}) *ErrorCode {
	return &ErrorCode{
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
	}
}

func ParseErrorCode(str string) (ec *ErrorCode, err error) {
	err = xjson.UnmarshalFromString(str, &ec)

	return
}

func (ce *ErrorCode) Code() int {
	return ce.ErrorCode
}

func (ce *ErrorCode) Msg() string {
	return ce.Message
}

func (ce *ErrorCode) ToData() interface{} {
	return ce.Data
}

func (ce *ErrorCode) Error() (str string) {
	if data, err := xjson.Marshal(ce); nil != err {
		return
	} else {
		str = string(data)
	}

	return
}
