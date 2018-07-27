package threadsafe

import (
	"sync"
	"time"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type CounterPerDuration struct {
	counter *base.CounterPerDuration
	mutex   *sync.Mutex
}

func NewCounterPerDuration(d time.Duration) *CounterPerDuration {
	return &CounterPerDuration{
		counter: base.NewCounterPerDuration(d),
		mutex:   &sync.Mutex{},
	}
}

func (c *CounterPerDuration) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

func (c *CounterPerDuration) Report() (int64, time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}
