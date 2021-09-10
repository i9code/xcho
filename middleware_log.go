package xcho

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/i9code/xlog"
	"github.com/i9code/xutil/xjson"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valyala/fasttemplate"
)

type (
	MiddlewareLoggerConfig struct {
		// 确定是不是要走中间件
		skipper middleware.Skipper
		// 格式
		Format string `yaml:"format"`

		pool *sync.Pool

		template *fasttemplate.Template
	}
)

var (
	defaultLoggerConfig = MiddlewareLoggerConfig{
		skipper: middleware.DefaultSkipper,
		Format: `"remote_ip":"${remote_ip}",` +
			`"status":${status},` +
			`"method":"${method}","uri":"${uri}",` +
			`"error":"${error}",` +
			`"latency_human":"${latency_human}",` +
			`"out":${out}}`,
	}
)

func logFunc(config MiddlewareLoggerConfig) echo.MiddlewareFunc {
	if config.skipper == nil {
		config.skipper = defaultLoggerConfig.skipper
	}
	if config.Format == "" {
		config.Format = defaultLoggerConfig.Format
	}

	config.pool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	config.template = fasttemplate.New(config.Format, "${", "}")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			buf := config.pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer config.pool.Put(buf)

			if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "remote_ip":
					return buf.WriteString(c.RealIP())
				case "uri":
					return buf.WriteString(req.RequestURI)
				case "method":
					return buf.WriteString(req.Method)
				case "path":
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					return buf.WriteString(p)
				case "status":
					n := res.Status
					if n >= 300 {
						buf.WriteString("\x1b[1;31m")
						buf.WriteString(strconv.Itoa(n))
						return buf.WriteString("\x1b[0m")
					}

					return buf.WriteString(strconv.Itoa(n))
				case "error":
					if err != nil {
						// Error may contain invalid JSON e.g. `"`
						b, _ := xjson.Marshal(err.Error())
						b = b[1 : len(b)-1]
						buf.WriteString("\x1b[1;31m")
						buf.Write(b)
						return buf.WriteString("\x1b[0m")
					}
				case "latency":
					l := stop.Sub(start)
					return buf.WriteString(strconv.FormatInt(int64(l), 10))
				case "latency_human":
					return buf.WriteString(stop.Sub(start).String())
				case "out":
					out, _ := xjson.Marshal(res)
					return buf.WriteString(string(out))
				default:
					switch {
					case strings.HasPrefix(tag, "header:"):
						return buf.Write([]byte(c.Request().Header.Get(tag[7:])))
					case strings.HasPrefix(tag, "query:"):
						return buf.Write([]byte(c.QueryParam(tag[6:])))
					case strings.HasPrefix(tag, "form:"):
						return buf.Write([]byte(c.FormValue(tag[5:])))
					case strings.HasPrefix(tag, "cookie:"):
						cookie, err := c.Cookie(tag[7:])
						if err == nil {
							return buf.Write([]byte(cookie.Value))
						}
					}
				}

				return 0, nil
			}); err != nil {
				return
			}

			xlog.Infof(buf.String())

			return
		}
	}
}
