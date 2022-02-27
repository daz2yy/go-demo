package mango

import (
	"fmt"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (router *router) AddRoute(method string, pattern string, f HandlerFunc) {
	url := method + "-" + pattern
	router.handlers[url] = f
}

func (router *router) GET(pattern string, f HandlerFunc) {
	router.AddRoute("GET", pattern, f)
}

func (router *router) POST(pattern string, f HandlerFunc) {
	router.AddRoute("POST", pattern, f)
}

func (router *router) handle(ctx *Context) {
	url := ctx.Method + "-" + ctx.Path
	if handler, ok := router.handlers[url]; ok {
		handler(ctx)
	} else {
		fmt.Fprintf(ctx.Writer, "404 Not Found: %q\n", ctx.Path)
	}
}
