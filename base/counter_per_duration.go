package base

import (
	"time"
)

type CounterPerDuration struct {
	previousCounter *Counter
	currentCounter  *Counter
	duration        time.Duration
	lastUpdate      time.Time
}

func NewCounterPerDuration(d time.Duration) *CounterPerDuration {
	return &CounterPerDuration{
		previousCounter: NewCounter(),
		currentCounter:  NewCounter(),
		duration:        d,
		lastUpdate:      time.Now(),
	}
}

func (c *CounterPerDuration) maybeShift() {
	if time.Since(c.lastUpdate) >= c.duration {
		c.lastUpdate = time.Now()
		*c.previousCounter = *c.currentCounter
		c.currentCounter = NewCounter()
	}
}

func (c *CounterPerDuration) Mark() {
	c.maybeShift()
	c.currentCounter.Mark()
}

func (c *CounterPerDuration) Report() (int64, time.Time) {
	c.maybeShift()
	return c.previousCounter.Report(), c.lastUpdate
}

func (c *CounterPerDuration) Reset() {
	c.previousCounter.Reset()
	c.currentCounter.Reset()
}
