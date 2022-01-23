package result

import (
	"github.com/Bibob7/reqCon/pkg/engine"
	"time"
)

type reqResult struct {
	success  bool
	duration time.Duration
	response engine.Response
}

func NewReqResult(success bool, duration time.Duration, response engine.Response) engine.ReqResult {
	return &reqResult{success: success, duration: duration, response: response}
}

func (r reqResult) Success() bool {
	return r.success
}

func (r reqResult) Duration() time.Duration {
	return r.duration
}

func (r reqResult) Response() engine.Response {
	return r.response
}
