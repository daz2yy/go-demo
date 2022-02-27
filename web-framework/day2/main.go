package main

// 基于 day1/base3 的内容做扩展
// Context 作用：
// 1. DRY 应用：每次请求返回 response 其实都需要设置 header，返回值格式化，通过 context 统一处理
// 2. 中间值保存：动态路由的 :param 参数存放，中间件处理结果的存放等
// Context 的生命周期：
// 1. 随着请求到来而诞生
// 2. 随着请求结束而销毁
// 所以，Context 上下文，就是一个辅助类，帮助我们处理这次请求过程中处理一些通用的功能，附带存储功能，让业务处理过程尽可能简单

import (
	"net/http"

	"mango"
)

func main() {
	engine := mango.New()

	// engine.GET("/", func(rw http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(rw, "URL.Path = %q", req.URL.Path)
	// })

	// engine.POST("/", func(rw http.ResponseWriter, req *http.Request) {
	// 	for k, v := range req.Header {
	// 		fmt.Fprintf(rw, "Header[%q] = [%q]", k, v)
	// 	}
	// })

	engine.GET("/", func(ctx *mango.Context) {
		ctx.String(http.StatusOK, "URL.Path = %q", ctx.Path)
	})

	engine.GET("/hello", func(ctx *mango.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, "hello %s, you're at path: %s", name, ctx.Path)
	})

	engine.POST("/header", func(ctx *mango.Context) {
		value := ctx.PostForm("world")
		ctx.JSON(http.StatusOK, mango.H{
			"hello": value,
		})
	})

	engine.PATCH("/patch", func(ctx *mango.Context) {
		value := ctx.PostForm("world")
		ctx.HTML(http.StatusOK, "<html><dir><h1>html</h1></dir></html>"+value)
	})

	engine.Run(":8899")
}
