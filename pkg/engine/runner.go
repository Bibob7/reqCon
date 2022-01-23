package engine

type Runner struct {
	Client        Client
	ResultHandler ResultHandler
}

func NewRunner(client Client, resultHandler ResultHandler) Runner {
	return Runner{Client: client, ResultHandler: resultHandler}
}

func (r Runner) Run(c Config) error {
	errChan := make(chan error, 1)
	reqResultChan := make(chan ReqResult, c.ReqNum)
	resultChan := r.ResultHandler.Handle(reqResultChan, c.ReqNum)
	go r.concurrentRun(c, reqResultChan, errChan)
	for {
		select {
		case err := <-errChan:
			return err
		case result := <-resultChan:
			result.Render()
			return nil
		default:
			// do nothing
		}
	}
}

// concurrentRun sends concurrent requests to the target,
// but the number of concurrent requests is limited by the config
func (r Runner) concurrentRun(c Config, reqResultChan chan ReqResult, errChan chan error) {
	maxConcurrent := make(chan struct{}, c.ConcNum)

	for i := 0; i < c.ReqNum; i++ {
		maxConcurrent <- struct{}{}
		go func(req Request, reqTime chan ReqResult, errChan chan error) {
			err := r.Client.SendReq(req, reqResultChan)
			if err != nil {
				errChan <- err
			}
			<-maxConcurrent
		}(c.Request, reqResultChan, errChan)
	}
}
