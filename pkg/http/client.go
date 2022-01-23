package http

import (
	"bytes"
	"github.com/Bibob7/reqCon/pkg/engine"
	"github.com/Bibob7/reqCon/pkg/result"
	"net/http"
	"time"
)

type Client struct{}

func (c *Client) SendReq(req engine.Request, reqResults chan engine.ReqResult) error {
	success := true

	buffer := &bytes.Buffer{}
	_, err := buffer.Read(req.Body())
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest(req.Method(), req.URL(), buffer)
	if err != nil {
		return err
	}
	client := http.Client{}
	startTime := time.Now()
	res, err := client.Do(httpReq)
	if err != nil {
		success = false
	}
	resStatusCode := res.StatusCode
	endTime := time.Now()
	reqTime := endTime.Sub(startTime)
	defer res.Body.Close()
	var bodyContent []byte
	res.Body.Read(bodyContent)
	reqResults <- result.NewReqResult(success, reqTime, NewResponse(nil, resStatusCode, nil))
	return nil
}
