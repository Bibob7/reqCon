package engine

import (
	"time"
)

type Client interface {
	SendReq(req Request, reqResults chan ReqResult) error
}

type Request interface {
	Method() string
	URL() string
	Header() map[string]string
	Body() []byte
}

type Response interface {
	Header() map[string]string
	StatusCode() int
	Body() []byte
}

type ResultHandler interface {
	Handle(reqResultChan chan ReqResult, reqNum int) chan Result
}

type ReqResult interface {
	Success() bool
	Duration() time.Duration
	Response() Response
}

type Result interface {
	Render()
}
