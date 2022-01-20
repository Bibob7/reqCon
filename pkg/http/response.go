package http

type Response struct {
	Header     map[string]string
	StatusCode int
	Body       []byte
}
