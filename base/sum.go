package base

type Sum struct {
	n   int64
	sum float64
}

func NewSum() *Sum {
	return &Sum{
		n:   0,
		sum: 0.0,
	}
}

func (s *Sum) Mark(val float64) {
	s.sum += val
	s.n++
}

func (s *Sum) Report() (float64, int64) {
	return s.sum, s.n
}

func (s *Sum) Reset() {
	s.n = 0
	s.sum = 0.0
}
