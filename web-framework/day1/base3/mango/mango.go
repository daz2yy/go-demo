package mango

import (
	"fmt"
	"log"
	"net/http"
)

type Mango struct {
	router map[string]http.HandlerFunc
}

func New() *Mango {
	return &Mango{
		router: make(map[string]http.HandlerFunc),
	}
}

func (engine *Mango) AddRoute(method string, pattern string, f http.HandlerFunc) {
	url := method + "-" + pattern
	engine.router[url] = f
}

func (engine *Mango) GET(pattern string, f http.HandlerFunc) {
	engine.AddRoute("GET", pattern, f)
}

func (engine *Mango) POST(pattern string, f http.HandlerFunc) {
	engine.AddRoute("POST", pattern, f)
}

func (engine *Mango) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	url := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[url]; ok {
		handler(rw, req)
	} else {
		fmt.Fprintf(rw, "404 Not Found: %q\n", req.URL)
	}
}

func (engine *Mango) Run(addr string) {
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(addr, engine))
}
