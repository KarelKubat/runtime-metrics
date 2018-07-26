package runtimemetrics

import (
	"time"
)

type sumPerDuration struct {
	previousSum *sum
	currentSum  *sum
	duration    time.Duration
	lastUpdate  time.Time
}

func newSumPerDuration(d time.Duration) *sumPerDuration {
	return &sumPerDuration{
		previousSum: newSum(),
		currentSum:  newSum(),
		duration:    d,
		lastUpdate:  time.Now(),
	}
}

func (s *sumPerDuration) maybeShift() {
	if time.Since(s.lastUpdate) >= s.duration {
		s.lastUpdate = time.Now()
		*s.previousSum = *s.currentSum
		s.currentSum = newSum()
	}
}

func (s *sumPerDuration) mark(val float64) {
	s.maybeShift()
	s.currentSum.mark(val)
}

func (s *sumPerDuration) report() (float64, int64, time.Time) {
	s.maybeShift()
	sum, n := s.previousSum.report()
	return sum, n, s.lastUpdate
}
