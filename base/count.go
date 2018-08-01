package base

import "sync"

type Count struct {
	counter int64
	mutex   *sync.Mutex
}

// NewCount returns a reference to this metric type.
func NewCount() *Count {
	return &Count{
		counter: 0,
		mutex:   &sync.Mutex{},
	}
}

// Mark marks the occurence of a "tick".
func (c *Count) Mark() {
	if c != nil {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		c.counter++
	}
}

// Report returns the number of observed "ticks".
func (c *Count) Report() int64 {
	if c == nil {
		return 0
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter
}

// Reset resets the metric.
func (c *Count) Reset() {
	if c != nil {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		c.counter = 0
	}
}
