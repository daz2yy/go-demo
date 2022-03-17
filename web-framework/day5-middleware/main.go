package main

// 中间件
// 1. 让用户能实现自定义插件
// 2. 支持不同路由组使用不同的插件
// 3. 支持更新数据到 Context

import (
	"log"
	"net/http"
	"time"

	"mango"
)

func v2Middleware() mango.HandlerFunc {
	return func(c *mango.Context) {
		t := time.Now()
		c.Fail(500, "Internal server error")
		log.Printf("[%d] [%s] in [%v] for groupV2", c.StatusCode, c.Request.URL.Path, time.Since(t))
	}
}

func main() {
	engine := mango.New()
	// 添加全局中间件
	engine.Use(mango.Logger())
	engine.GET("/", func(ctx *mango.Context) {
		ctx.String(http.StatusOK, "URL.Path = %q", ctx.Path)
	})

	v1 := engine.Group("/v2")
	v1.Use(v2Middleware())
	v1.GET("/hello", func(ctx *mango.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, "hello %s, you're at path: %s", name, ctx.Path)
	})

	engine.Run(":8899")
}
