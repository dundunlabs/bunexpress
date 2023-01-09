package bunexpress

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type Group struct {
	group *bunrouter.Group
}

func (g Group) NewGroup(path string, opts ...bunrouter.GroupOption) *Group {
	return &Group{group: g.group.NewGroup(path, opts...)}
}

func (g Group) WithMiddleware(middleware MiddlewareFunc) *Group {
	return &Group{group: g.group.WithMiddleware(BunMiddlewareFunc(middleware))}
}

func (g Group) WithGroup(path string, fn func(g *Group)) {
	fn(g.NewGroup(path))
}

func (g Group) Handle(method string, path string, handler HandlerFunc) {
	g.group.Handle(method, path, BunHandlerFunc(handler))
}

func (g Group) GET(path string, handler HandlerFunc) {
	g.Handle(http.MethodGet, path, handler)
}

func (g Group) POST(path string, handler HandlerFunc) {
	g.Handle("POST", path, handler)
}

func (g Group) PUT(path string, handler HandlerFunc) {
	g.Handle("PUT", path, handler)
}

func (g Group) DELETE(path string, handler HandlerFunc) {
	g.Handle("DELETE", path, handler)
}

func (g Group) PATCH(path string, handler HandlerFunc) {
	g.Handle("PATCH", path, handler)
}

func (g Group) HEAD(path string, handler HandlerFunc) {
	g.Handle("HEAD", path, handler)
}

func (g Group) OPTIONS(path string, handler HandlerFunc) {
	g.Handle("OPTIONS", path, handler)
}

func (g Group) LoadRoutes(routes []Route) {
	for _, route := range routes {
		sg := &g
		if route.Middleware != nil {
			sg = sg.WithMiddleware(route.Middleware)
		}
		if len(route.Children) > 0 {
			sg = sg.NewGroup(route.Path, route.Options...)
			sg.LoadRoutes(route.Children)
		} else {
			sg.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
