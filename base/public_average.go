package runtimemetrics

type Average struct {
	summer *Sum
}

func NewAverage() *Average {
	return &Average{
		summer: NewSum(),
	}
}

func (a *Average) Mark(val float64) {
	a.summer.Mark(val)
}

func (a *Average) Report() (float64, int64) {
	sum, n := a.summer.Report()
	if n == 0 {
		return 0.0, n
	}
	return sum / float64(n), n
}
