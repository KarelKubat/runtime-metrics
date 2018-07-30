package threadsafe

import (
	"sync"
	"time"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type SumPerDuration struct {
	sum   *base.SumPerDuration
	mutex *sync.Mutex
}

// NewSumPerDuration returns a reference to an initialized threadsafe.SumPerDuration.
// The argument is the time window for summing.
func NewSumPerDuration(d time.Duration) *SumPerDuration {
	return &SumPerDuration{
		sum:   base.NewSumPerDuration(d),
		mutex: &sync.Mutex{},
	}
}

// Mark registers a float64 observation.
func (s *SumPerDuration) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum.Mark(val)
}

// Report returns the sum, number of cases and end time.
func (s *SumPerDuration) Report() (float64, int64, time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.sum.Report()
}

// Reset resets the metric.
func (s *SumPerDuration) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum.Reset()
}
