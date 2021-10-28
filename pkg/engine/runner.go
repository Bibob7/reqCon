package engine

import (
	http2 "github.com/Bibob7/reqCon/pkg/http"
	"net/http"
)

func Run(url string, con int, num int) {
	queue := make(chan bool, con)

	for i := 0; i < num; i++ {
		queue <- true
		go func(url string) {
			err := http2.SendReq(http2.Request{
				http.MethodGet,
				url,
				map[string]string{},
				[]byte{},
			})
			if err != nil {
				// failed request
			}
		}(url)
	}

	for i := 0; i < con; i++ {
		queue <- true
	}
}
