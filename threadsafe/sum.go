package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Sum struct {
	summer *base.Sum
	mutex  *sync.Mutex
}

// NewSum returns a reference to an initialized threadsafe.Sum.
func NewSum() *Sum {
	return &Sum{
		summer: base.NewSum(),
		mutex:  &sync.Mutex{},
	}
}

// Mark registers a float64 observation.
func (s *Sum) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.summer.Mark(val)
}

// Report returns the sum of observations and the number of cases.
func (s *Sum) Report() (float64, int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.summer.Report()
}

// Reset resets the metric.
func (s *Sum) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.summer.Reset()
}
