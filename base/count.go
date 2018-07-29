package base

type Count struct {
	counter int64
}

// NewCount returns a reference to this metric type.
func NewCount() *Count {
	return &Count{
		counter: 0,
	}
}

// Mark marks the occurence of a "tick".
func (c *Count) Mark() {
	c.counter++
}

// Report returns the number of observed "ticks".
func (c *Count) Report() int64 {
	return c.counter
}

// Reset resets the metric.
func (c *Count) Reset() {
	c.counter = 0
}
