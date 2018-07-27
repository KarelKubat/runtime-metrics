package named

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type CounterPerDuration struct {
	handler *threadsafe.CounterPerDuration
	name    string
}

func NewCounterPerDuration(n string, d time.Duration) (*CounterPerDuration, error) {
	metric := &CounterPerDuration{
		handler: threadsafe.NewCounterPerDuration(d),
		name:    n,
	}
	return metric, reg.registerCounterPerDuration(metric)
}

func (c *CounterPerDuration) Mark() {
	c.handler.Mark()
}

func (c *CounterPerDuration) Report() (int64, time.Time) {
	return c.handler.Report()
}
