package runtimemetrics

import (
	"sync"
)

type Sum struct {
	summer *sum
	mutex  *sync.Mutex
}

func NewSum() *Sum {
	return &Sum{
		summer: newSum(),
		mutex:  &sync.Mutex{},
	}
}

func (s *Sum) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.summer.mark(val)
}

func (s *Sum) Report() (float64, int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.summer.report()
}
