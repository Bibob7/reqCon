package engine

import (
	"github.com/Bibob7/reqCon/pkg/engine/config"
	http2 "github.com/Bibob7/reqCon/pkg/http"
)

func Run(c config.Config) error {
	queue := make(chan bool, c.ConcNum)

	for i := 0; i < c.ReqNum; i++ {
		queue <- true
		go func(url string) {
			err := http2.SendReq(c.Request)
			if err != nil {
				// failed request
			}
		}(c.URL)
	}

	for i := 0; i < c.ConcNum; i++ {
		queue <- true
	}
	return nil
}
