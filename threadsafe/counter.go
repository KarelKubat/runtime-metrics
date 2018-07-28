package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Counter struct {
	counter *base.Counter
	mutex   *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		counter: base.NewCounter(),
		mutex:   &sync.Mutex{},
	}
}

func (c *Counter) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

func (c *Counter) Report() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}

func (c *Counter) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Reset()
}

