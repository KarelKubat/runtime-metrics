package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Average struct {
	avg   *base.Average
	mutex *sync.Mutex
}

func NewAverage() *Average {
	return &Average{
		avg:   base.NewAverage(),
		mutex: &sync.Mutex{},
	}
}

func (a *Average) Mark(val float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.avg.Mark(val)
}

func (a *Average) Report() (float64, int64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.avg.Report()
}

func (a *Average) Reset() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.Reset()
}
