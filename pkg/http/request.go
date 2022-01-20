package http

type Request struct {
	Method string            `yaml:"method" json:"method"`
	Url    string            `yaml:"url" json:"url"`
	Header map[string]string `yaml:"header" json:"header"`
	Body   string            `yaml:"body" json:"body"`
}
