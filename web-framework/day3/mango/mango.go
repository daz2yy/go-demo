package mango

import (
	"log"
	"net/http"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

type (
	Mango struct {
		*RouterGroup
		router *router
		groups []*RouterGroup // store all groups
	}

	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Mango // 为了方便在 Engine 层用户可以直接使用引擎里的 router 添加路由，加在这
	}
)

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func New() *Mango {
	engine := &Mango{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

// todo: 不是 / 结尾
func (group *RouterGroup) AddRoute(method string, pattern string, handle HandlerFunc) {
	fullPattern := group.prefix + pattern
	group.engine.router.AddRoute(method, fullPattern, handle)
}

func (group *RouterGroup) GET(pattern string, handle HandlerFunc) {
	group.AddRoute("GET", pattern, handle)
}

func (group *RouterGroup) POST(pattern string, handle HandlerFunc) {
	group.AddRoute("POST", pattern, handle)
}

func (group *RouterGroup) PATCH(pattern string, handle HandlerFunc) {
	group.AddRoute("PATCH", pattern, handle)
}

func (engine *Mango) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := newContext(rw, req)
	engine.router.handle(ctx)
}

func (engine *Mango) Run(addr string) {
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(addr, engine))
}
