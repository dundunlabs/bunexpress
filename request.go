package bunexpress

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/uptrace/bunrouter"
)

type Body []byte

func (body Body) Bind(v any) error {
	return json.Unmarshal(body, v)
}

func (body Body) BindX(v any) {
	if err := body.Bind(v); err != nil {
		panic(err)
	}
}

func (body Body) Map() (m map[string]any, err error) {
	err = body.Bind(&m)
	return
}

func (body Body) MapX() map[string]any {
	m, err := body.Map()
	if err != nil {
		panic(err)
	}
	return m
}

type Request struct {
	req  bunrouter.Request
	body Body
}

func NewRequest(req bunrouter.Request) *Request {
	return &Request{req: req}
}

func (r *Request) Method() string {
	return r.req.Method
}

func (r *Request) Route() string {
	return r.req.Route()
}

func (r *Request) Params() bunrouter.Params {
	return r.req.Params()
}

func (r *Request) URL() *url.URL {
	return r.req.URL
}

func (r *Request) Query() map[string]any {
	var m map[string]any
	for k, v := range r.URL().Query() {
		if len(v) > 1 {
			m[k] = v
		} else {
			m[k] = v[0]
		}
	}

	return m
}

func (r *Request) Body() (Body, error) {
	if r.body == nil {
		b, err := io.ReadAll(r.req.Body)
		if err != nil {
			return nil, err
		}
		r.body = b
	}

	return r.body, nil
}

func (r *Request) BodyX() Body {
	body, err := r.Body()
	if err != nil {
		panic(err)
	}
	return body
}
