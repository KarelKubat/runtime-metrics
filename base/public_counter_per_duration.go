package runtimemetrics

import (
	"sync"
	"time"
)

type CounterPerDuration struct {
	counter *counterPerDuration
	mutex   *sync.Mutex
}

func NewCounterPerDuration(d time.Duration) *CounterPerDuration {
	return &CounterPerDuration{
		counter: newCounterPerDuration(d),
		mutex:   &sync.Mutex{},
	}
}

func (c *CounterPerDuration) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.mark()
}

func (c *CounterPerDuration) Report() (int64, time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.report()
}
