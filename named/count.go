package named

import "github.com/KarelKubat/runtime-metrics/threadsafe"

type Count struct {
	handler *threadsafe.Count
	name    string
}

func NewCount(n string) *Count {
	return &Count{
		handler: threadsafe.NewCount(),
		name:    n,
	}
}

func (c *Count) Mark() {
	c.handler.Mark()
}

func (c *Count) Report() int64 {
	return c.handler.Report()
}

func (c *Count) Reset() {
	c.handler.Reset()
}

func (c *Count) Name() string {
	return c.name
}
