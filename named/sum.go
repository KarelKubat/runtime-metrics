package named

import "github.com/KarelKubat/runtime-metrics/threadsafe"

type Sum struct {
	handler *threadsafe.Sum
	name    string
}

func NewSum(n string) (*Sum, error) {
	metric := &Sum{
		handler: threadsafe.NewSum(),
		name:    n,
	}
	return metric, reg.registerSum(metric)
}

func (c *Sum) Mark(val float64) {
	c.handler.Mark(val)
}

func (c *Sum) Report() (float64, int64) {
	return c.handler.Report()
}
