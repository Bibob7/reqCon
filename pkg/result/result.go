package result

import (
	"fmt"
	"time"
)

type Result struct {
	AVGDuration  time.Duration
	NumFailedReq int
	StatusCodes  []int
	MessageList  []string
}

func (r Result) Render() {
	fmt.Printf("Average Time: %v \n", r.AVGDuration)
	fmt.Printf("Failed Requests: %v \n", r.NumFailedReq)
	for _, message := range r.MessageList {
		fmt.Printf("%s\n", message)
	}
}
