package named

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type CountPerDuration struct {
	handler *threadsafe.CountPerDuration
	name    string
}

func NewCountPerDuration(n string, d time.Duration) *CountPerDuration {
	return &CountPerDuration{
		handler: threadsafe.NewCountPerDuration(d),
		name:    n,
	}
}

func (c *CountPerDuration) Mark() {
	c.handler.Mark()
}

func (c *CountPerDuration) Report() (int64, time.Time) {
	return c.handler.Report()
}

func (c *CountPerDuration) Reset() {
	c.handler.Reset()
}

func (c *CountPerDuration) Name() string {
	return c.name
}
