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
	g.Get("/hello", func(ctx *xcho.Context) (err error) {

		return ctx.JSON(http.StatusOK, "abc")
	}).Name = "测试"
}
