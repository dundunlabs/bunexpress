package bunexpress

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HandlerFunc func(req Request, res Response)

func BunHandlerFunc(handler HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		handler(Request{req}, Response{w})
		return nil
	}
}

func FromBunHandlerFunc(handler bunrouter.HandlerFunc) HandlerFunc {
	return func(req Request, res Response) {
		handler(res, req.Request)
	}
}
