package runtimemetrics

type sum struct {
	n   int64
	sum float64
}

func newSum() *sum {
	return &sum{
		n:   0,
		sum: 0.0,
	}
}

func (s *sum) mark(val float64) {
	s.sum += val
	s.n++
}

func (s *sum) report() (float64, int64) {
	return s.sum, s.n
}
