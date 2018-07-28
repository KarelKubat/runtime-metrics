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

func NewSumPerDuration(d time.Duration) *SumPerDuration {
	return &SumPerDuration{
		sum:   base.NewSumPerDuration(d),
		mutex: &sync.Mutex{},
	}
}

func (s *SumPerDuration) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum.Mark(val)
}

func (s *SumPerDuration) Report() (float64, int64, time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.sum.Report()
}

func (s *SumPerDuration) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum.Reset()
}

