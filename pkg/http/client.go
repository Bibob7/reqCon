package http

import (
	"bytes"
	"net/http"
	"time"
)

func SendReq(req Request, reqTimes chan time.Duration) error {
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
	startTime := time.Now()
	res, err := client.Do(httpReq)
	endTime := time.Now()
	reqTimes <- endTime.Sub(startTime)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var bodyContent []byte
	res.Body.Read(bodyContent)
	return nil
}
