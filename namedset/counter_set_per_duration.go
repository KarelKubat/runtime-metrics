package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type CounterPerDurationSet struct {
	set map[string]*named.CounterPerDuration
}

func NewCounterPerDurationSet() *CounterPerDurationSet {
	return &CounterPerDurationSet{
		set: map[string]*named.CounterPerDuration{},
	}
}

func (set *CounterPerDurationSet) Add(a *named.CounterPerDuration) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("CounterPerDuration %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *CounterPerDurationSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CounterPerDurationSet) Get(name string) (*named.CounterPerDuration, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("CounterPerDuration %q not in set", name)
	}
	return ret, nil
}
