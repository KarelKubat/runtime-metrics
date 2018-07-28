package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type CounterSet struct {
	set map[string]*named.Counter
}

func NewCounterSet() *CounterSet {
	return &CounterSet{
		set: map[string]*named.Counter{},
	}
}

func (set *CounterSet) Add(a *named.Counter) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("Counter %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *CounterSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CounterSet) Get(name string) (*named.Counter, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("Counter %q not in set", name)
	}
	return ret, nil
}
