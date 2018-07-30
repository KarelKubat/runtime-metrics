package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Count struct {
	counter *base.Count
	mutex   *sync.Mutex
}

// Newcount returns a reference to an initialized threadsafe.Count.
func NewCount() *Count {
	return &Count{
		counter: base.NewCount(),
		mutex:   &sync.Mutex{},
	}
}

// Mark registers an observation (a "tick") and internally increments the count.
func (c *Count) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

// Report returns the number of observations.
func (c *Count) Report() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}

// Reset resets the metric.
func (c *Count) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Reset()
}
