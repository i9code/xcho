package xcho

import (
	`encoding/json`
	`encoding/xml`
	`github.com/labstack/echo/v4`
	`github.com/vmihailenco/msgpack/v5`
	`google.golang.org/protobuf/proto`
)

const (
	// HttpMethodGet GET方法
	HttpMethodGet HttpMethod = "GET"
	// HttpMethodPost POST方法
	HttpMethodPost HttpMethod = "POST"
	// HttpMethodPut PUT方法
	HttpMethodPut HttpMethod = "PUT"
	// HttpMethodDelete DELETE方法
	HttpMethodDelete HttpMethod = "DELETE"
	// HttpMethodPatch PATCH方法
	HttpMethodPatch HttpMethod = "PATCH"
	// HttpMethodHead HEAD方法
	HttpMethodHead HttpMethod = "HEAD"
	// HttpMethodOptions OPTIONS方法
	HttpMethodOptions HttpMethod = "OPTIONS"
)

func data(ctx echo.Context, rsp interface{}, options *httpOptions) (err error) {
	var blob []byte
	var contentType string
	switch options.serializer {
	case serializerJson:
		contentType = echo.MIMEApplicationJSONCharsetUTF8
		blob, err = json.Marshal(rsp)
	case serializerXml:
		contentType = echo.MIMEApplicationXMLCharsetUTF8
		blob, err = xml.Marshal(rsp)
	case serializerMsgpack:
		contentType = echo.MIMEApplicationMsgpack
		blob, err = msgpack.Marshal(rsp)
	case serializerProto:
		contentType = echo.MIMEApplicationProtobuf
		blob, err = proto.Marshal(rsp.(proto.Message))
	case serializerBytes:
		contentType = options.contentType
		switch rsp.(type) {
		case []byte:
			blob = rsp.([]byte)
		case *[]byte:
			blob = *rsp.(*[]byte)
		}
	}
	if nil != err {
		return
	}

	// 写入数据
	ctx.Response().Header().Set(echo.HeaderContentType, contentType)
	// ctx.Response().Status = options.code
	ctx.Response().WriteHeader(options.code)
	_, err = ctx.Response().Write(blob)

	return
}
