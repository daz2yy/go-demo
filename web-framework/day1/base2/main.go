package main

// 统一路由处理入口，后续可以增加统一的逻辑处理，比如监控，日志，异常处理等

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(rw, "URL.Path = %q", req.URL.Path)
	case "/header":
		for k, v := range req.Header {
			fmt.Fprintf(rw, "Header[%q] = [%q]", k, v)
		}
	}
}

func main() {
	engine := new(Engine)

	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":8899", engine))
}
