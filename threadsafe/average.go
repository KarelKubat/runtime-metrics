package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Average struct {
	avg   *base.Average
	mutex *sync.Mutex
}

// NewAverage returns a reference to an initialized threadsafe.Average.
func NewAverage() *Average {
	return &Average{
		avg:   base.NewAverage(),
		mutex: &sync.Mutex{},
	}
}

// Mark registers a float64 observation.
func (a *Average) Mark(val float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Mark(val)
}

// Report returns the average over all observations and the number of cases.
func (a *Average) Report() (float64, int64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.avg.Report()
}

// Reset resets the metric.
func (a *Average) Reset() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.Reset()
}
