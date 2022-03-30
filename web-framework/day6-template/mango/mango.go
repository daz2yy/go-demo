package mango

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

type (
	Mango struct {
		*RouterGroup
		router        *router
		groups        []*RouterGroup     // store all groups
		htmlTemplates *template.Template // for html render
		funcMap       template.FuncMap   // fro html render
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

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		filePath := c.Param("filepath")
		if _, err := fs.Open(filePath); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// static router
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	pattern := path.Join(relativePath, "/*filepath")
	group.GET(pattern, handler)
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
	group.engine.router.addRoute(method, fullPattern, handle)
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

// add middleware
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Mango) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Mango) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (engine *Mango) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var middleware []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middleware = append(middleware, group.middlewares...)
		}
	}
	ctx := newContext(rw, req)
	ctx.handlers = middleware
	ctx.engine = engine
	engine.router.handle(ctx)
}

func (engine *Mango) Run(addr string) {
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(addr, engine))
}
