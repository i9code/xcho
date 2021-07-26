package xcho

type (
	HandlerFunc func(ctx *Context) (err error)
)
