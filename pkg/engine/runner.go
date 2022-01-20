package engine

import (
	"fmt"
	"github.com/Bibob7/reqCon/pkg/engine/config"
	http "github.com/Bibob7/reqCon/pkg/http"
	"time"
)

func Run(c config.Config) error {
	reqTimes := make(chan time.Duration, c.ReqNum)
	queue := make(chan bool, c.ConcNum)
	req, err := c.Request()
	if err != nil {
		return err
	}

	go func(reqTimes chan time.Duration) {
		var reqCounter int64
		var summedTimes int64
		for {
			duration := <-reqTimes
			reqCounter++
			summedTimes += duration.Nanoseconds()
			fmt.Printf("Average Time: %v \n", time.Duration(summedTimes/reqCounter))
		}
	}(reqTimes)

	for i := 0; i < c.ReqNum; i++ {
		queue <- true
		go func(req http.Request, reqTime chan time.Duration) {
			err := http.SendReq(req, reqTimes)
			if err != nil {
				fmt.Printf("Request failed: %v \n", err)
			}
			<-queue
		}(*req, reqTimes)
	}

	for i := 0; i < c.ConcNum; i++ {
		queue <- true
	}
	return nil
}
