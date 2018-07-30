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

// NewAveragePerDuration returns a reference to an initialized threadsafe.AveragePerDuration.
// The argument is the time window, such as 10*time.Second.
func NewAveragePerDuration(d time.Duration) *AveragePerDuration {
	return &AveragePerDuration{
		avg:   base.NewAveragePerDuration(d),
		mutex: &sync.Mutex{},
	}
}

// Mark registers a float64 observation.
func (a *AveragePerDuration) Mark(val float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Mark(val)
}

// Report returns the average over all observations, the number of cases, and the
// ending period for this average.
func (a *AveragePerDuration) Report() (float64, int64, time.Time) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.avg.Report()
}

// Reset resets the metric.
func (a *AveragePerDuration) Reset() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Reset()
}
