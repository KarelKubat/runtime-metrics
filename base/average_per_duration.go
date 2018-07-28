package base

import (
	"time"
)

type AveragePerDuration struct {
	summer *SumPerDuration
}

func NewAveragePerDuration(d time.Duration) *AveragePerDuration {
	return &AveragePerDuration{
		summer: NewSumPerDuration(d),
	}
}

func (a *AveragePerDuration) Mark(val float64) {
	a.summer.Mark(val)
}

func (a *AveragePerDuration) Report() (float64, int64, time.Time) {
	sum, n, stamp := a.summer.Report()
	if n == 0 {
		return 0.0, 0, stamp
	}
	return sum / float64(n), n, stamp
}

func (a *AveragePerDuration) Reset() {
	a.summer.Reset()
}