package threadsafe

import (
	"sync"

	base "github.com/KarelKubat/runtime-metrics/base"
)

type Sum struct {
	summer *base.Sum
	mutex  *sync.Mutex
}

func NewSum() *Sum {
	return &Sum{
		summer: base.NewSum(),
		mutex:  &sync.Mutex{},
	}
}

func (s *Sum) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.summer.Mark(val)
}

func (s *Sum) Report() (float64, int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.summer.Report()
}
