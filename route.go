package bunexpress

import "github.com/uptrace/bunrouter"

type Route struct {
	Path       string
	Method     string
	Handler    HandlerFunc
	Middleware MiddlewareFunc
	Children   []Route
	Options    []bunrouter.GroupOption
}
