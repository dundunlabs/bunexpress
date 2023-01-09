package bunexpress

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/uptrace/bunrouter"
)

type Router struct {
	*bunrouter.Router
	*Group
}

func New(router *bunrouter.Router) *Router {
	return &Router{
		Router: router,
		Group:  &Group{group: &router.Group},
	}
}

type DefaultHandler struct {
	router *Router
}

func (router *Router) DefaultHandler() DefaultHandler {
	handler := DefaultHandler{router}
	return handler
}

func (h DefaultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
	}()

	h.router.ServeHTTP(w, req)
}
