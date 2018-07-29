package base

import (
	"time"
)

type CountPerDuration struct {
	previousCount *Count
	currentCount  *Count
	duration      time.Duration
	lastUpdate    time.Time
}

// NewCountPerDuration returns a reference to this metric.
func NewCountPerDuration(d time.Duration) *CountPerDuration {
	return &CountPerDuration{
		previousCount: NewCount(),
		currentCount:  NewCount(),
		duration:      d,
		lastUpdate:    time.Now(),
	}
}

func (c *CountPerDuration) maybeShift() {
	if time.Since(c.lastUpdate) >= c.duration {
		c.lastUpdate = time.Now()
		*c.previousCount = *c.currentCount
		c.currentCount = NewCount()
	}
}

// Mark marks the occurrence of a "tick".
func (c *CountPerDuration) Mark() {
	c.maybeShift()
	c.currentCount.Mark()
}

// Report returns the number of observed "ticks" and the time until which the count was
// maintained.
func (c *CountPerDuration) Report() (int64, time.Time) {
	c.maybeShift()
	return c.previousCount.Report(), c.lastUpdate
}

// Reset resets the metric.
func (c *CountPerDuration) Reset() {
	c.previousCount.Reset()
	c.currentCount.Reset()
}
