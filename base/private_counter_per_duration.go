package runtimemetrics

import (
	"time"
)

type counterPerDuration struct {
	previousCounter *counter
	currentCounter  *counter
	duration        time.Duration
	lastUpdate      time.Time
}

func newCounterPerDuration(d time.Duration) *counterPerDuration {
	return &counterPerDuration{
		previousCounter: newCounter(),
		currentCounter:  newCounter(),
		duration:        d,
		lastUpdate:      time.Now(),
	}
}

func (c *counterPerDuration) maybeShift() {
	if time.Since(c.lastUpdate) >= c.duration {
		c.lastUpdate = time.Now()
		*c.previousCounter = *c.currentCounter
		c.currentCounter = newCounter()
	}
}

func (c *counterPerDuration) mark() {
	c.maybeShift()
	c.currentCounter.mark()
}

func (c *counterPerDuration) report() (int64, time.Time) {
	c.maybeShift()
	return c.previousCounter.report(), c.lastUpdate
}
