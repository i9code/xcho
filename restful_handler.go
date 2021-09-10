package xcho

type restfulHandler func(ctx *Context) (rsp interface{}, err error)
