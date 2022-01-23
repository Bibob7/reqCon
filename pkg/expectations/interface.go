package expectations

import "github.com/Bibob7/reqCon/pkg/engine"

type Result struct {
	Success bool
	ErrMsg  string
}

type Expectation interface {
	Validate(res engine.Response) Result
}
