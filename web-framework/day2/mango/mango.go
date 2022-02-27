package mango

import (
	"log"
	"net/http"
)

type Mango struct {
	router *router
}

// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

func New() *Mango {
	return &Mango{
		router: newRouter(),
	}
}

func (engine *Mango) AddRoute(method string, pattern string, handle HandlerFunc) {
	engine.router.AddRoute(method, pattern, handle)
}

func (engine *Mango) GET(pattern string, handle HandlerFunc) {
	engine.AddRoute("GET", pattern, handle)
}

func (engine *Mango) POST(pattern string, handle HandlerFunc) {
	engine.AddRoute("POST", pattern, handle)
}

func (engine *Mango) PATCH(pattern string, handle HandlerFunc) {
	engine.AddRoute("PATCH", pattern, handle)
}

func (engine *Mango) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := newContext(rw, req)
	engine.router.handle(ctx)
}

func (engine *Mango) Run(addr string) {
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(addr, engine))
}
