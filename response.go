package bunexpress

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type Response struct {
	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

func (res *Response) Status(statusCode int) *Response {
	res.w.WriteHeader(statusCode)
	return res
}

func (res *Response) Json(v any) error {
	return bunrouter.JSON(res.w, v)
}

func (res *Response) JsonX(v any) {
	if err := res.Json(v); err != nil {
		panic(err)
	}
}
