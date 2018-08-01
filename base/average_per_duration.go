package base

import (
	"time"
)

type AveragePerDuration struct {
	summer *SumPerDuration
}

// NewAveragePerDuration returns a reference to this metric type.
func NewAveragePerDuration(d time.Duration) *AveragePerDuration {
	return &AveragePerDuration{
		summer: NewSumPerDuration(d),
	}
}

// Mark marks an observation of a floating point value.
func (a *AveragePerDuration) Mark(val float64) {
	if a != nil && a.summer != nil {
		a.summer.Mark(val)
	}
}

// Report returns the average, number of observed values, and time until which
// the avarage was computed. The observation started at the returned timestamp
// minus the duration.
func (a *AveragePerDuration) Report() (float64, int64, time.Time) {
	if a == nil || a.summer == nil {
		return 0.0, 0, time.Now()
	}
	sum, n, stamp := a.summer.Report()
	if n == 0 {
		return 0.0, 0, stamp
	}
	return sum / float64(n), n, stamp
}

// Reset resets the metric.
func (a *AveragePerDuration) Reset() {
	if a != nil && a.summer != nil {
		a.summer.Reset()
	}
}
