package named

import "github.com/KarelKubat/runtime-metrics/threadsafe"

type Counter struct {
	handler *threadsafe.Counter
	name    string
}

func NewCounter(n string) *Counter {
	return &Counter{
		handler: threadsafe.NewCounter(),
		name:    n,
	}
}

func (c *Counter) Mark() {
	c.handler.Mark()
}

func (c *Counter) Report() int64 {
	return c.handler.Report()
}

func (c *Counter) Reset() {
	c.handler.Reset()
}

func (c *Counter) Name() string {
	return c.name
}
