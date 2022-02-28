package main

// 增加路由分组设置
// 1. 本质上是一种便捷代码设置的功能，分组管理，支持嵌套分组
// 2. 放在 Engine 里，用户可以直接使用
// 3. 分组除了关联相同前缀，还可以管理中间件

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
