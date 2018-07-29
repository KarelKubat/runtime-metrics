package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Count struct {
	counter *base.Count
	mutex   *sync.Mutex
}

func NewCount() *Count {
	return &Count{
		counter: base.NewCount(),
		mutex:   &sync.Mutex{},
	}
}

func (c *Count) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Mark()
}

func (c *Count) Report() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.Report()
}

func (c *Count) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.Reset()
}

