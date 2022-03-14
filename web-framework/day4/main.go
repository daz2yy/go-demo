package main

// 动态路由解析
// 1. 支持路由：/v1/hello/:name, 可以匹配 /v1/hello/abc, /v1/hello/asd 等
// 2. 支持路由：/v1/file/*, *可以匹配任意长度的文件路径，常用于前端文件匹配；支持开头是 *
// 3. 参数解析保存
// 4. 前缀树实现

import (
	"net/http"

	"mango"
)

func main() {
	engine := mango.New()
	engine.GET("/", func(ctx *mango.Context) {
		ctx.String(http.StatusOK, "URL.Path = %q", ctx.Path)
	})

	v1 := engine.Group("/v1")
	v1.GET("/hello", func(ctx *mango.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, "hello %s, you're at path: %s", name, ctx.Path)
	})

	v1.POST("/header", func(ctx *mango.Context) {
		value := ctx.PostForm("world")
		ctx.JSON(http.StatusOK, mango.H{
			"hello": value,
		})
	})

	v1.PATCH("/patch", func(ctx *mango.Context) {
		value := ctx.PostForm("world")
		ctx.HTML(http.StatusOK, "<html><dir><h1>html</h1></dir></html>"+value)
	})

	engine.Run(":8899")
}
