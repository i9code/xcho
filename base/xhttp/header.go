package xhttp

import (
	"fmt"
	"net/url"
)

const (
	// HeaderAcceptLanguage 可接受的语言
	HeaderAcceptLanguage      = "Accept-Language"
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-IP"
	HeaderXRequestID          = "X-Request-ID"
	HeaderXRequestedWith      = "X-Requested-With"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"

	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

	HeaderStrictTransportSecurity         = "Strict-Transport-Security"
	HeaderXContentTypeOptions             = "X-Content-Type-Options"
	HeaderXXSSProtection                  = "X-XSS-Protection"
	HeaderXFrameOptions                   = "X-Frame-Options"
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	HeaderXCSRFToken                      = "X-CSRF-Token"
	HeaderReferrerPolicy                  = "Referrer-Policy"
)

const (
	// ContentDispositionTypeAttachment 附件下载
	ContentDispositionTypeAttachment ContentDispositionType = "attachment"
	// ContentDispositionTypeInline 浏览器直接打开
	ContentDispositionTypeInline ContentDispositionType = "inline"
)

const (
	// HttpParameterTypeHeader 请求头
	HttpParameterTypeHeader HttpParameterType = "header"
	// HttpParameterTypePathParameter 路径参数
	HttpParameterTypePathParameter HttpParameterType = "path"
)

const (
	// AuthTypeBasic 基本Http授权验证
	AuthTypeBasic AuthType = "basic"
	// AuthTypeToken 基本传Token的授权验证
	AuthTypeToken AuthType = "token"
)

// AuthType 授权类型
type AuthType string

type (
	// HttpMethod Http方法
	HttpMethod string

	// ContentDispositionType 下载类型
	ContentDispositionType string

	// HttpParameterType Http额外参数类型
	HttpParameterType string

	// HttpParameter Http额外参数接口
	HttpParameter interface {
		// Type 类型
		Type() HttpParameterType
		// Key 键
		Key() string
		// Value 值
		Value() string
	}
)

// ContentDisposition 解决附件下载乱码
func ContentDisposition(filename string, dispositionType ContentDispositionType) (disposition string) {
	// 文件名需要编码
	filename = url.QueryEscape(filename)
	disposition = fmt.Sprintf("%s; filename=%s;filename*=utf-8''%s", dispositionType, filename, filename)

	return
}
