package named

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type SumPerDuration struct {
	handler *threadsafe.SumPerDuration
	name    string
}

func NewSumPerDuration(n string, d time.Duration) (*SumPerDuration, error) {
	metric := &SumPerDuration{
		handler: threadsafe.NewSumPerDuration(d),
		name:    n,
	}
	return metric, reg.registerSumPerDuration(metric)
}

func (c *SumPerDuration) Mark(val float64) {
	c.handler.Mark(val)
}

func (c *SumPerDuration) Report() (float64, int64, time.Time) {
	return c.handler.Report()
}
