package expectations

type Expectation interface {
	Validate() bool
}
