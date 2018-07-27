package named

import "fmt"

type registry struct {
	average            []*Average
	averagePerDuration []*AveragePerDuration
	counter            []*Counter
	counterPerDuration []*CounterPerDuration
	sum                []*Sum
	sumPerDuration     []*SumPerDuration
}

func newRegistry() *registry {
	return &registry{
		average:            []*Average{},
		averagePerDuration: []*AveragePerDuration{},
		counter:            []*Counter{},
		counterPerDuration: []*CounterPerDuration{},
		sum:                []*Sum{},
		sumPerDuration:     []*SumPerDuration{},
	}
}

var reg *registry

func init() {
	reg = newRegistry()
}

func (reg *registry) checkName(n, kind string) error {
	for _, handler := range reg.average {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.averagePerDuration {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.counter {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.counterPerDuration {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.sum {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	for _, handler := range reg.sumPerDuration {
		if handler.name == n {
			return fmt.Errorf("duplicate %s name %v", kind, n)
		}
	}
	return nil
}

func (reg *registry) registerAverage(avg *Average) error {
	if err := reg.checkName(avg.name, "Average"); err != nil {
		return err
	}
	reg.average = append(reg.average, avg)
	return nil
}

func (reg *registry) registerAveragePerDuration(avg *AveragePerDuration) error {
	if err := reg.checkName(avg.name, "AveragePerDuration"); err != nil {
		return err
	}
	reg.averagePerDuration = append(reg.averagePerDuration, avg)
	return nil
}

func (reg *registry) registerCounter(avg *Counter) error {
	if err := reg.checkName(avg.name, "Counter"); err != nil {
		return err
	}
	reg.counter = append(reg.counter, avg)
	return nil
}

func (reg *registry) registerCounterPerDuration(avg *CounterPerDuration) error {
	if err := reg.checkName(avg.name, "CounterPerDuration"); err != nil {
		return err
	}
	reg.counterPerDuration = append(reg.counterPerDuration, avg)
	return nil
}

func (reg *registry) registerSum(avg *Sum) error {
	if err := reg.checkName(avg.name, "Sum"); err != nil {
		return err
	}
	reg.sum = append(reg.sum, avg)
	return nil
}

func (reg *registry) registerSumPerDuration(avg *SumPerDuration) error {
	if err := reg.checkName(avg.name, "SumPerDuration"); err != nil {
		return err
	}
	reg.sumPerDuration = append(reg.sumPerDuration, avg)
	return nil
}
