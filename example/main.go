package main

import (
	"fmt"
	"net/http"

	"github.com/dundunlabs/bunexpress"
	"github.com/uptrace/bunrouter"
)

func errorMiddleware(next bunexpress.HandlerFunc) bunexpress.HandlerFunc {
	return func(req bunexpress.Request, res bunexpress.Response) {
		defer func() {
			if err := recover(); err != nil {
				res.Status(http.StatusInternalServerError).
					JSON(bunrouter.H{"message": fmt.Sprint(err)})
			}
		}()

		next(req, res)
	}
}

func pingHandler(req bunexpress.Request, res bunexpress.Response) {
	res.JSON(bunrouter.H{"message": "pong"})
}

func helloHandler(req bunexpress.Request, res bunexpress.Response) {
	res.JSON(bunrouter.H{"message": "Hello world!"})
}

func errorHandler(req bunexpress.Request, res bunexpress.Response) {
	panic("Oops!")
}

var routes = []bunexpress.Route{
	{
		Path:    "/ping",
		Method:  http.MethodGet,
		Handler: pingHandler,
	},
	{
		Path:       "/api",
		Middleware: errorMiddleware,
		Children: []bunexpress.Route{
			{
				Path:    "/hello",
				Method:  http.MethodGet,
				Handler: helloHandler,
			},
			{
				Path:    "/error",
				Method:  http.MethodGet,
				Handler: errorHandler,
			},
		},
	},
}

func main() {
	br := bunrouter.New()
	router := bunexpress.New(br)

	router.LoadRoutes(routes)

	handler := router.DefaultHandler()
	http.ListenAndServe(":8080", handler)
}
