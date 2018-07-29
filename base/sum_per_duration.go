package base

import (
	"time"
)

type SumPerDuration struct {
	previousSum *Sum
	currentSum  *Sum
	duration    time.Duration
	lastUpdate  time.Time
}

// NewSumPerDuration returns a reference to this metric type.
func NewSumPerDuration(d time.Duration) *SumPerDuration {
	return &SumPerDuration{
		previousSum: NewSum(),
		currentSum:  NewSum(),
		duration:    d,
		lastUpdate:  time.Now(),
	}
}

func (s *SumPerDuration) maybeShift() {
	if time.Since(s.lastUpdate) >= s.duration {
		s.lastUpdate = time.Now()
		*s.previousSum = *s.currentSum
		s.currentSum = NewSum()
	}
}

// Mark marks the occurrence of a floating point value.
func (s *SumPerDuration) Mark(val float64) {
	s.maybeShift()
	s.currentSum.Mark(val)
}

// Report returns the sum of the observed values, the number of observed values, and the time
// until which the sum was computed.
func (s *SumPerDuration) Report() (float64, int64, time.Time) {
	s.maybeShift()
	sum, n := s.previousSum.Report()
	return sum, n, s.lastUpdate
}

// Reset resets this metric.
func (s *SumPerDuration) Reset() {
	s.previousSum.Reset()
	s.currentSum.Reset()
}
