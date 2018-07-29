package base

type Sum struct {
	n   int64
	sum float64
}

// NewSum returns a reference to this metric type.
func NewSum() *Sum {
	return &Sum{
		n:   0,
		sum: 0.0,
	}
}

// Mark adds the occurrence of a floating point value.
func (s *Sum) Mark(val float64) {
	s.sum += val
	s.n++
}

// Report returns the sum and number of observed values.
func (s *Sum) Report() (float64, int64) {
	return s.sum, s.n
}

// Reset resets the metric.
func (s *Sum) Reset() {
	s.n = 0
	s.sum = 0.0
}
