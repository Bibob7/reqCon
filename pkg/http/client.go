package http

import (
	"bytes"
	"net/http"
)

func SendReq(req Request) error {
	buffer := &bytes.Buffer{}
	_, err := buffer.Read(req.Body)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest(req.Method, req.Url, buffer)
	if err != nil {
		return err
	}
	client := http.Client{}
	_, err = client.Do(httpReq)
	if err != nil {
		return err
	}
	return nil
}
