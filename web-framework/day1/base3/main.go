package main

// 1. 修改 mod 指向相对路径
// 2. 用 New 的方式创建, 初始化 map
// 3. GET, POST 方法封装
// 4. Run 启动 http 服务
// 5. handler, ok := router[r]

import (
	"fmt"
	"net/http"

	"mango"
)

func main() {
	engine := mango.New()

	engine.GET("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "URL.Path = %q", req.URL.Path)
	})

	engine.POST("/", func(rw http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(rw, "Header[%q] = [%q]", k, v)
		}
	})

	engine.Run(":8899")
}
