package runtimemetrics

import (
	"sync"
	"time"
)

type SumPerDuration struct {
	sum   *sumPerDuration
	mutex *sync.Mutex
}

func NewSumPerDuration(d time.Duration) *SumPerDuration {
	return &SumPerDuration{
		sum:   newSumPerDuration(d),
		mutex: &sync.Mutex{},
	}
}

func (s *SumPerDuration) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum.mark(val)
}

func (s *SumPerDuration) Report() (float64, int64, time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.sum.report()
}
