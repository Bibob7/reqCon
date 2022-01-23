package config

import (
	"github.com/Bibob7/reqCon/pkg/expectations"
	"io/ioutil"
)

type Config struct {
	ReqNum          int               `yaml:"number" json:"number"`
	ConcNum         int               `yaml:"concurrent" json:"concurrent"`
	RequestConfig   Request           `yaml:"request" json:"request"`
	ExpectationList expectations.List `yaml:"expectations" json:"expectations"`
}

type Request struct {
	Method string            `yaml:"method" json:"method"`
	Url    string            `yaml:"url" json:"url"`
	Header map[string]string `yaml:"header,omitempty" json:"header,omitempty"`
	Body   Body              `yaml:"body,omitempty" json:"body,omitempty"`
}

type Body struct {
	String   string `yaml:"string,omitempty" json:"string,omitempty"`
	FilePath string `yaml:"path,omitempty" json:"path,omitempty"`
}

func (c *Config) Body() ([]byte, error) {
	if c.RequestConfig.Body.FilePath != "" {
		return ioutil.ReadFile(c.RequestConfig.Body.FilePath)
	}
	return []byte(c.RequestConfig.Body.String), nil
}
