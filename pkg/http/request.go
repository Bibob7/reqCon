package http

type Request struct {
	Method string
	Url    string
	Header map[string]string
	Body   []byte
}
