package xcho

import (
	"context"
	"os"
	"os/signal"

	"github.com/i9code/xutils/valid"
	"github.com/i9code/xutils/xhttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type initFunc func(echo *Echo)

type Echo struct {
	*echo.Echo

	options *options
}

// 新建Echo服务
func New(opts ...option) (server *Echo) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	// 创建Echo服务器
	server = &Echo{
		Echo: echo.New(),
	}

	server.options = options

	server.HideBanner = !options.banner

	// 初始化
	for _, init := range options.inits {
		init(server)
	}

	// 处理路由
	if 0 != len(options.routes) {
		group := &Group{group: server.Group(options.context)}
		for _, route := range options.routes {
			route(group)
		}
	}

	// 数据验证
	if options.validate {
		server.Validator = valid.NewValidator()
	}

	// 初始化绑定
	if options.binder {
		server.Binder = &binder{}
	}

	// 处理错误
	server.HTTPErrorHandler = echo.HTTPErrorHandler(options.error)

	// 初始化中间件
	server.Pre(middleware.MethodOverride())
	server.Pre(middleware.RemoveTrailingSlash())

	// server.Use(middleware.CSRF())
	server.Use(logFunc(defaultLoggerConfig))
	server.Use(middleware.RequestID())
	// 配置跨域
	if options.crosEnable {
		cors := middleware.DefaultCORSConfig
		cors.AllowMethods = append(cors.AllowMethods, string(xhttp.HttpMethodOptions))
		cors.AllowOrigins = options.cros.origins
		cors.AllowCredentials = options.cros.credentials
		server.Use(middleware.CORSWithConfig(cors))
	}

	// 增加内置中间件
	// 增加Jwt
	if options.jwtEnable {
		server.Use(jwtFunc(options.jwt))
	}
	// 增加Http签名验证
	if options.signatureEnable {
		server.Use(signatureFunc(options.signature))
	}
	// 增加权限验证
	if options.casbinEnable {
		server.Use(casbinFunc(options.casbin))
	}
	// 打印堆栈信息
	// 方便调试，默认处理没有换行，很难内眼查看堆栈信息
	server.Use(panicStackFunc(options.panicStack))

	// 增加自定义上下文
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&Context{
				Context: c,
				Jwt:     options.jwt,
			})
		}
	})

	return
}

func (e *Echo) Start() (err error) {
	// 在另外的协程中启动服务器，实现优雅地关闭（Graceful Shutdown）
	go func() {
		err = e.Echo.Start(e.options.addr)
	}()

	// 等待系统退出中断并响应
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), e.options.shutdownTimeout)
	defer cancel()

	e.Shutdown(ctx)

	return
}
