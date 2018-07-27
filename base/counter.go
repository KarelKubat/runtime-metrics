package base

type Counter struct {
	counter int64
}

func NewCounter() *Counter {
	return &Counter{
		counter: 0,
	}
}

func (c *Counter) Mark() {
	c.counter++
}

func (c *Counter) Report() int64 {
	return c.counter
}
