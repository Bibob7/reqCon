package result

import (
	"github.com/Bibob7/reqCon/pkg/engine"
	"github.com/Bibob7/reqCon/pkg/expectations"
	"github.com/schollz/progressbar/v3"
	"time"
)

type resultHandler struct {
	expList expectations.List
}

func NewResultHandler(expList expectations.List) engine.ResultHandler {
	return &resultHandler{expList: expList}
}

func (r resultHandler) Handle(reqResultChan chan engine.ReqResult, reqNum int) chan engine.Result {
	resultChan := make(chan engine.Result, 1)
	go r.listenToResults(reqResultChan, reqNum, resultChan)
	return resultChan
}

func (r resultHandler) listenToResults(reqResultChan chan engine.ReqResult, reqNum int, resultChan chan engine.Result) {
	pBar := progressbar.Default(int64(reqNum))
	var summedTimes int64
	var failedReq int
	for i := 1; i <= reqNum; i++ {
		reqResult := <-reqResultChan
		pBar.Add(1)
		summedTimes += reqResult.Duration().Nanoseconds()
		if !reqResult.Success() {
			failedReq++
		}
	}
	resultChan <- Result{
		AVGDuration:  time.Duration(summedTimes / int64(reqNum)),
		NumFailedReq: failedReq,
	}
}
