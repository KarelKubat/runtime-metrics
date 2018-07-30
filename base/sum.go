package base

import "sync"

type Sum struct {
	n     int64
	sum   float64
	mutex *sync.Mutex
}

// NewSum returns a reference to this metric type.
func NewSum() *Sum {
	return &Sum{
		n:     0,
		sum:   0.0,
		mutex: &sync.Mutex{},
	}
}

// Mark adds the occurrence of a floating point value.
func (s *Sum) Mark(val float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sum += val
	s.n++
}

// Report returns the sum and number of observed values.
func (s *Sum) Report() (float64, int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.sum, s.n
}

// Reset resets the metric.
func (s *Sum) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.n = 0
	s.sum = 0.0
}
