package bunexpress

import (
	"github.com/uptrace/bunrouter"
)

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

func BunMiddlewareFunc(mdw MiddlewareFunc) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return BunHandlerFunc(mdw(FromBunHandlerFunc(next)))
	}
}
