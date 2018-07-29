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

func NewCountPerDuration(d time.Duration) *CountPerDuration {
	return &CountPerDuration{
		counter: base.NewCountPerDuration(d),
		mutex:   &sync.Mutex{},
	}
}

func (c *CountPerDuration) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

func (c *CountPerDuration) Report() (int64, time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}

func (c *CountPerDuration) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Reset()
}
