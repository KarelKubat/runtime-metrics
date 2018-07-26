package runtimemetrics

type counter struct {
	counter int64
}

func newCounter() *counter {
	return &counter{
		counter: 0,
	}
}

func (c *counter) mark() {
	c.counter++
}

func (c *counter) report() int64 {
	return c.counter
}
