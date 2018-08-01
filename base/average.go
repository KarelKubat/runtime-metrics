package base

// Average is the metric type for averages.
type Average struct {
	summer *Sum
}

// NewAverage returns a reference to this metric type.
func NewAverage() *Average {
	return &Average{
		summer: NewSum(),
	}
}

// Mark marks the occurrence of a floating point value.
func (a *Average) Mark(val float64) {
	if a != nil && a.summer != nil {
		a.summer.Mark(val)
	}
}

// Report returns the average and number of observed values.
func (a *Average) Report() (float64, int64) {
	if a == nil || a.summer == nil {
		return 0.0, 0
	}
	sum, n := a.summer.Report()
	if n == 0 {
		return 0.0, 0
	}
	return sum / float64(n), n
}

// Reset resets the metric.
func (a *Average) Reset() {
	if a != nil && a.summer != nil {
		a.summer.Reset()
	}
}
