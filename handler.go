package bunexpress

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HandlerFunc func(req *Request, res *Response)

func BunHandlerFunc(handler HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		handler(NewRequest(req), NewResponse(w))
		return nil
	}
}

func FromBunHandlerFunc(handler bunrouter.HandlerFunc) HandlerFunc {
	return func(req *Request, res *Response) {
		handler(res.w, req.req)
	}
}
