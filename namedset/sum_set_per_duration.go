package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type SumPerDurationSet struct {
	set map[string]*named.SumPerDuration
}

func NewSumPerDurationSet() *SumPerDurationSet {
	return &SumPerDurationSet{
		set: map[string]*named.SumPerDuration{},
	}
}

func (set *SumPerDurationSet) Add(a *named.SumPerDuration) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("SumPerDuration %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *SumPerDurationSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *SumPerDurationSet) Get(name string) (*named.SumPerDuration, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("SumPerDuration %q not in set", name)
	}
	return ret, nil
}
