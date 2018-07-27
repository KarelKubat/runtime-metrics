package named

import "github.com/KarelKubat/runtime-metrics/threadsafe"

type Counter struct {
	handler *threadsafe.Counter
	name    string
}

func NewCounter(n string) (*Counter, error) {
	metric := &Counter{
		handler: threadsafe.NewCounter(),
		name:    n,
	}
	return metric, reg.registerCounter(metric)
}

func (c *Counter) Mark() {
	c.handler.Mark()
}

func (c *Counter) Report() int64 {
	return c.handler.Report()
}
