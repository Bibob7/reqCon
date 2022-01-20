package config

import "github.com/Bibob7/reqCon/pkg/http"

type Config struct {
	URL     string       `yaml:"url" json:"url"`
	ReqNum  int          `yaml:"number" json:"number"`
	ConcNum int          `yaml:"concurrent" json:"concurrent"`
	Request http.Request `yaml:"request" json:"request"`
}
