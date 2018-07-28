package threadsafe

import (
	"sync"
	"time"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type AveragePerDuration struct {
	avg   *base.AveragePerDuration
	mutex *sync.Mutex
}

func NewAveragePerDuration(d time.Duration) *AveragePerDuration {
	return &AveragePerDuration{
		avg:   base.NewAveragePerDuration(d),
		mutex: &sync.Mutex{},
	}
}

func (a *AveragePerDuration) Mark(val float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Mark(val)
}

func (a *AveragePerDuration) Report() (float64, int64, time.Time) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.avg.Report()
}

func (a *AveragePerDuration) Reset() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Reset()
}
