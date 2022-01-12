package main

import (
	"github.com/i9code/xcho/base/xhttp"
	"net/http"

	"github.com/i9code/xcho"
)

func main() {
	server := xcho.New()

	server.Start(xcho.Routes(func(group *xcho.Group) {
		apiMount(group.Group("/api"))

	}))
}

func apiMount(g *xcho.Group, mfs ...xcho.MiddlewareFunc) {
	g.Post("/test", func(ctx *xcho.Context) (err error) {

		return ctx.JSON(http.StatusOK, "post")
	}).Name = "测试"

	g.Get("/test", func(ctx *xcho.Context) (err error) {

		return ctx.JSON(http.StatusOK, "get")
	}).Name = "测试"

	g.Get("/http/client", httpClient).Name = "测试http client"
}

func httpClient(ctx *xcho.Context) (err error) {
	httpClient, _ := xhttp.NewHttpClient(&xhttp.HttpClientConfig{})
	rsp, err := httpClient.R().Get("https://cart.jd.com/gate.action?pid=10386048547&pcount=1&ptype=1")
	if nil != err {
		//	xlog.Debug(err)
	}
	//xlog.Debug(rsp)

	return ctx.JSON(http.StatusOK, rsp.Body())
}
