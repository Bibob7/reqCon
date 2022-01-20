package expectations

import (
	"github.com/Bibob7/reqCon/pkg/http"
)

type HeaderExists struct {
	ExpHeaders []string `yaml:"headers"`
}

func (e HeaderExists) Validate(res http.Response) bool {
	for _, expHeader := range e.ExpHeaders {
		if _, ok := res.Header[expHeader]; !ok {
			return false
		}
	}
	return true
}

type HeaderWithValueExists struct {
	ExpHeadersWithValues map[string]string `yaml:"headers"`
}

func (e HeaderWithValueExists) Validate(res http.Response) bool {
	for expHeader, expValue := range e.ExpHeadersWithValues {
		v, ok := res.Header[expHeader]
		if !ok {
			return false
		}
		if v != expValue {
			return false
		}
	}
	return true
}
