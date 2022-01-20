package expectations

import (
	"github.com/Bibob7/reqCon/pkg/http"
	"strings"
)

type BodyEqual struct {
	Expected string `yaml:"exp"`
}

func (e BodyEqual) Validate(res http.Response) bool {
	return e.Expected == string(res.Body)
}

type BodyContains struct {
	Expected string `yaml:"exp"`
}

func (e BodyContains) Validate(res http.Response) bool {
	return strings.Contains(string(res.Body), e.Expected)
}
