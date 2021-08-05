package main

import (
	"net/http"

	"github.com/i9code/xcho"
)

func main() {
	server := xcho.New(xcho.Routes(func(group *xcho.Group) {
		apiMount(group.Group("/api"))
	}))

	server.Start()
}

func apiMount(g *xcho.Group, mfs ...xcho.MiddlewareFunc) {
	g.GET("/hello", func(ctx *xcho.Context) (err error) {

		return ctx.JSON(http.StatusOK, "abc")
	}).Name = "测试"
}
