package main

import (
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
}
