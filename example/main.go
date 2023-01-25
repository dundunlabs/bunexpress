package main

import (
	"fmt"
	"net/http"

	"github.com/dundunlabs/bunexpress"
	"github.com/uptrace/bunrouter"
)

func errorMiddleware(next bunexpress.HandlerFunc) bunexpress.HandlerFunc {
	return func(req *bunexpress.Request, res *bunexpress.Response) {
		defer func() {
			if err := recover(); err != nil {
				res.Status(http.StatusInternalServerError).
					JsonX(bunrouter.H{"message": fmt.Sprint(err)})
			}
		}()

		next(req, res)
	}
}

func pingHandler(req *bunexpress.Request, res *bunexpress.Response) {
	res.JsonX(bunrouter.H{"message": "pong"})
}

type HelloBody struct {
	Name string `json:"name"`
}

func helloHandler(req *bunexpress.Request, res *bunexpress.Response) {
	m := req.BodyX().MapX()
	fmt.Println("map:", m)

	var body HelloBody
	req.BodyX().BindX(&body)

	res.JsonX(bunrouter.H{"message": fmt.Sprintf("Hello %s!", body.Name)})
}

func errorHandler(req *bunexpress.Request, res *bunexpress.Response) {
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
				Method:  http.MethodPost,
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
