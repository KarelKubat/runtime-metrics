package named

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type AveragePerDuration struct {
	handler *threadsafe.AveragePerDuration
	name    string
}

func NewAveragePerDuration(n string, d time.Duration) *AveragePerDuration {
	return &AveragePerDuration{
		handler: threadsafe.NewAveragePerDuration(d),
		name:    n,
	}
}

func (a *AveragePerDuration) Mark(val float64) {
	a.handler.Mark(val)
}

func (a *AveragePerDuration) Report() (float64, int64, time.Time) {
	return a.handler.Report()
}

func (a *AveragePerDuration) Reset() {
	a.handler.Reset()
}

func (a *AveragePerDuration) Name() string {
	return a.name
}
