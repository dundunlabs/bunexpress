package bunexpress

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type Response struct {
	http.ResponseWriter
}

func (res Response) Status(statusCode int) Response {
	res.WriteHeader(statusCode)
	return res
}

func (res Response) JSON(v any) {
	if err := bunrouter.JSON(res, v); err != nil {
		panic(err)
	}
}
