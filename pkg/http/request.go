package http

import "github.com/Bibob7/reqCon/pkg/engine"

type request struct {
	method string
	url    string
	header map[string]string
	body   []byte
}

func NewRequest(method string, url string, header map[string]string, body []byte) engine.Request {
	return &request{method: method, url: url, header: header, body: body}
}

func (r request) Method() string {
	return r.method
}

func (r request) URL() string {
	return r.url
}

func (r request) Header() map[string]string {
	return r.header
}

func (r request) Body() []byte {
	return r.body
}
