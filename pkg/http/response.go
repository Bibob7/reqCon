package http

import "github.com/Bibob7/reqCon/pkg/engine"

type response struct {
	header     map[string]string
	statusCode int
	body       []byte
}

func NewResponse(header map[string]string, statusCode int, body []byte) engine.Response {
	return &response{header: header, statusCode: statusCode, body: body}
}

func (r response) Header() map[string]string {
	return r.header
}

func (r response) StatusCode() int {
	return r.statusCode
}

func (r response) Body() []byte {
	return r.body
}
