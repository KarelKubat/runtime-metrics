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

func (s *SumPerDuration) Mark(val float64) {
	s.maybeShift()
	s.currentSum.Mark(val)
}

func (s *SumPerDuration) Report() (float64, int64, time.Time) {
	s.maybeShift()
	sum, n := s.previousSum.Report()
	return sum, n, s.lastUpdate
}

func (s *SumPerDuration) Reset() {
	s.previousSum.Reset()
	s.currentSum.Reset()
}
