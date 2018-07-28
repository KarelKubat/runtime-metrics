package named

import (
	"fmt"
	named "github.com/KarelKubat/runtime-metrics/named"
)

type registry struct {
	average            []*named.Average
	averagePerDuration []*named.AveragePerDuration
	counter            []*named.Counter
	counterPerDuration []*named.CounterPerDuration
	sum                []*named.Sum
	sumPerDuration     []*named.SumPerDuration
}

var reg *registry

func init() {
	reg = &registry{
		average:            []*named.Average{},
		averagePerDuration: []*named.AveragePerDuration{},
		counter:            []*named.Counter{},
		counterPerDuration: []*named.CounterPerDuration{},
		sum:                []*named.Sum{},
		sumPerDuration:     []*named.SumPerDuration{},
	}
}

func (reg *registry) checkName(n, kind string) error {
	for _, handler := range reg.average {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.averagePerDuration {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.counter {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.counterPerDuration {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.sum {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.sumPerDuration {
		if handler.Name() == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	return nil
}

func RegisterAverage(avg *named.Average) error {
	if err := reg.checkName(avg.Name(), "Average"); err != nil {
		return err
	}
	reg.average = append(reg.average, avg)
	return nil
}

func RegisterAveragePerDuration(avg *named.AveragePerDuration) error {
	if err := reg.checkName(avg.Name(), "AveragePerDuration"); err != nil {
		return err
	}
	reg.averagePerDuration = append(reg.averagePerDuration, avg)
	return nil
}

func RegisterCounter(avg *named.Counter) error {
	if err := reg.checkName(avg.Name(), "Counter"); err != nil {
		return err
	}
	reg.counter = append(reg.counter, avg)
	return nil
}

func RegisterCounterPerDuration(avg *named.CounterPerDuration) error {
	if err := reg.checkName(avg.Name(), "CounterPerDuration"); err != nil {
		return err
	}
	reg.counterPerDuration = append(reg.counterPerDuration, avg)
	return nil
}

func RegisterSum(avg *named.Sum) error {
	if err := reg.checkName(avg.Name(), "Sum"); err != nil {
		return err
	}
	reg.sum = append(reg.sum, avg)
	return nil
}

func RegisterSumPerDuration(avg *named.SumPerDuration) error {
	if err := reg.checkName(avg.Name(), "SumPerDuration"); err != nil {
		return err
	}
	reg.sumPerDuration = append(reg.sumPerDuration, avg)
	return nil
}
