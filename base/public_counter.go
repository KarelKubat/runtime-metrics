package runtimemetrics

import (
	"sync"
)

type Counter struct {
	counter *counter
	mutex   *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		mutex:   &sync.Mutex{},
		counter: newCounter(),
	}
}

func (c *Counter) Mark() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter.mark()
}

func (c *Counter) Report() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.counter.report()
}
