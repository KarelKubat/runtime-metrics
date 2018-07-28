package named

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type CounterPerDuration struct {
	handler *threadsafe.CounterPerDuration
	name    string
}

func NewCounterPerDuration(n string, d time.Duration) *CounterPerDuration {
	return &CounterPerDuration{
		handler: threadsafe.NewCounterPerDuration(d),
		name:    n,
	}
}

func (c *CounterPerDuration) Mark() {
	c.handler.Mark()
}

func (c *CounterPerDuration) Report() (int64, time.Time) {
	return c.handler.Report()
}

func (c *CounterPerDuration) Reset() {
	c.handler.Reset()
}

func (c *CounterPerDuration) Name() string {
	return c.name
}
