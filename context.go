package xcho

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	`github.com/i9code/xutils/xhttp`
	json `github.com/json-iterator/go`
	"github.com/labstack/echo/v4"
)

const defaultIndent = "  "

// Context 自定义的Echo上下文
type Context struct {
	echo.Context

	// Jwt配置
	Jwt JwtConfig
}

func (c *Context) Bind(data interface{}) (err error) {
	if err = c.Context.Bind(data); nil != err {
		return
	}
	err = c.Validate(data)

	return
}

func (c *Context) Subject(subject interface{}) (err error) {
	var (
		token  string
		claims jwt.Claims
	)

	if token, err = c.Jwt.runExtractor(c.Context); nil != err {
		return
	}
	if claims, _, err = c.Jwt.Parse(token); nil != err {
		return
	}
	// 从Token中反序列化主题数据
	err = json.UnmarshalFromString(claims.(*jwt.StandardClaims).Subject, &subject)

	return
}

func (c *Context) JwtToken(domain string, data interface{}, expire time.Duration) (token string, id string, err error) {
	return c.Jwt.Token(domain, data, expire)
}

func (c *Context) HttpFile(file http.File) (err error) {
	defer func() {
		_ = file.Close()
	}()

	var info os.FileInfo
	if info, err = file.Stat(); nil != err {
		return
	}

	http.ServeContent(c.Response(), c.Request(), info.Name(), info.ModTime(), file)

	return
}

func (c *Context) HttpAttachment(file http.File, name string) error {
	return c.contentDisposition(file, name, xhttp.ContentDispositionTypeAttachment)
}

func (c *Context) HttpInline(file http.File, name string) error {
	return c.contentDisposition(file, name, xhttp.ContentDispositionTypeInline)
}

func (c *Context) contentDisposition(file http.File, name string, dispositionType xhttp.ContentDispositionType) error {
	c.Response().Header().Set(xhttp.HeaderContentDisposition, xhttp.ContentDisposition(name, dispositionType))

	return c.HttpFile(file)
}

func (c *Context) JSON(code int, i interface{}) (err error) {
	indent := ""
	if _, pretty := c.QueryParams()["pretty"]; c.Echo().Debug || pretty {
		indent = defaultIndent
	}
	return c.json(code, i, indent)
}

func (c *Context) JSONPretty(code int, i interface{}, indent string) (err error) {
	return c.json(code, i, indent)
}

func (c *Context) JSONBlob(code int, b []byte) (err error) {
	return c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
}

func (c *Context) JSONP(code int, callback string, i interface{}) (err error) {
	return c.jsonPBlob(code, callback, i)
}

func (c *Context) JSONPBlob(code int, callback string, b []byte) (err error) {
	c.writeContentType(echo.MIMEApplicationJavaScriptCharsetUTF8)
	c.Response().WriteHeader(code)
	if _, err = c.Response().Write([]byte(callback + "(")); err != nil {
		return
	}
	if _, err = c.Response().Write(b); err != nil {
		return
	}
	_, err = c.Response().Write([]byte(");"))

	return
}

func (c *Context) jsonPBlob(code int, callback string, i interface{}) (err error) {
	enc := json.NewEncoder(c.Response())
	_, pretty := c.QueryParams()["pretty"]
	if c.Echo().Debug || pretty {
		enc.SetIndent("", "  ")
	}
	c.writeContentType(echo.MIMEApplicationJavaScriptCharsetUTF8)
	c.Response().WriteHeader(code)
	if _, err = c.Response().Write([]byte(callback + "(")); err != nil {
		return
	}
	if err = enc.Encode(i); err != nil {
		return
	}
	if _, err = c.Response().Write([]byte(");")); err != nil {
		return
	}

	return
}

func (c *Context) json(code int, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if "" != indent {
		enc.SetIndent("", indent)
	}
	c.writeContentType(echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(code)

	return enc.Encode(i)
}

func (c *Context) writeContentType(value string) {
	header := c.Response().Header()
	if "" == header.Get(echo.HeaderContentType) {
		header.Set(echo.HeaderContentType, value)
	}
}
