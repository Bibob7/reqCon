package expectations

import (
	"fmt"
	"github.com/Bibob7/reqCon/pkg/engine"
	"strings"
)

type BodyEqual struct {
	Expected string `yaml:"exp"`
}

func (e BodyEqual) Validate(res engine.Response) Result {
	actualBody := string(res.Body())
	if e.Expected == actualBody {
		return Result{
			Success: true,
		}
	}
	return Result{
		Success: false,
		ErrMsg:  fmt.Sprintf("Body is not equal expected body. \n exptected: %s \n actual: %s \n", e.Expected, actualBody),
	}
}

type BodyContains struct {
	Expected string `yaml:"exp"`
}

func (e BodyContains) Validate(res engine.Response) Result {
	actualBody := string(res.Body())
	if strings.Contains(actualBody, e.Expected) {
		return Result{
			Success: true,
		}
	}

	return Result{
		Success: false,
		ErrMsg:  fmt.Sprintf("Body does not contains expected substring. \n exptected: %s \n actual: %s \n", e.Expected, actualBody),
	}
}
