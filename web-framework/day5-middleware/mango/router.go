package mango

// 动态路由解析
// 1. 支持路由：/v1/hello/:name, 可以匹配 /v1/hello/abc, /v1/hello/asd 等
// 2. 支持路由：/v1/file/*, *可以匹配任意长度的文件路径，常用于前端文件匹配；支持开头是 *
// 3. 参数解析保存
// 4. 前缀树实现

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析路由，用 map 来保存，方便后续处理
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, f HandlerFunc) {
	parts := parsePattern(pattern)

	url := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	log.Printf("AddRoute %4s - %s", method, pattern)
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[url] = f
}

func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 解析参数变量并调用函数
	params := make(map[string]string)
	node := root.search(searchParts, 0)
	if node != nil {
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}

	return nil, nil
}

func (r *router) GET(pattern string, f HandlerFunc) {
	r.addRoute("GET", pattern, f)
}

func (r *router) POST(pattern string, f HandlerFunc) {
	r.addRoute("POST", pattern, f)
}

// 处理接收到的请求
func (r *router) handle(ctx *Context) {
	node, params := r.getRoute(ctx.Method, ctx.Path)
	if node != nil {
		ctx.Params = params
		url := ctx.Method + "-" + node.pattern
		ctx.handlers = append(ctx.handlers, r.handlers[url])
	} else {
		ctx.handlers = append(ctx.handlers, func(c *Context) {
			ctx.String(http.StatusNotFound, "404 Not Found: %q\n", ctx.Path)
		})
	}
	ctx.Next()
}
