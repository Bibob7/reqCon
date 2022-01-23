package expectations

import (
	"fmt"
	"github.com/Bibob7/reqCon/pkg/engine"
)

type HeaderExists struct {
	ExpHeaders []string `yaml:"headers"`
}

func (e HeaderExists) Validate(res engine.Response) Result {
	for _, expHeader := range e.ExpHeaders {
		if _, ok := res.Header()[expHeader]; !ok {
			return Result{
				Success: false,
				ErrMsg:  fmt.Sprintf("Exptected header \"%s\" is missing", expHeader),
			}
		}
	}
	return Result{
		Success: true,
	}
}

type HeaderWithValueExists struct {
	ExpHeadersWithValues map[string]string `yaml:"headers"`
}

func (e HeaderWithValueExists) Validate(res engine.Response) Result {
	for expHeader, expValue := range e.ExpHeadersWithValues {
		v, ok := res.Header()[expHeader]
		if !ok {
			return Result{
				Success: false,
				ErrMsg:  fmt.Sprintf("Exptected header \"%s\" is missing", expHeader),
			}
		}
		if v != expValue {
			return Result{
				Success: false,
				ErrMsg:  fmt.Sprintf("Exptected header \"%s\" does not contains expected value. \n exptected: %s \n actual: %s \n", expHeader, expValue, v),
			}
		}
	}
	return Result{
		Success: true,
	}
}
