package threadsafe

import (
	"sync"
	"time"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type CountPerDuration struct {
	counter *base.CountPerDuration
	mutex   *sync.Mutex
}

// NewCountPerDuration returns a reference to an initialized threadsafe.CountPerDuration. The argument
// is the time window over which counting will apply.
func NewCountPerDuration(d time.Duration) *CountPerDuration {
	return &CountPerDuration{
		counter: base.NewCountPerDuration(d),
		mutex:   &sync.Mutex{},
	}
}

// Mark registers an observation.
func (c *CountPerDuration) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

// Report returns the number of observations and the ending time of counting.
func (c *CountPerDuration) Report() (int64, time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}

// Reset resets the metric.
func (c *CountPerDuration) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Reset()
}
