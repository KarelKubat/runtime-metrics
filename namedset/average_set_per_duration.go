package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type AveragePerDurationSet struct {
	set map[string]*named.AveragePerDuration
}

func NewAveragePerDurationSet() *AveragePerDurationSet {
	return &AveragePerDurationSet{
		set: map[string]*named.AveragePerDuration{},
	}
}

func (set *AveragePerDurationSet) Add(a *named.AveragePerDuration) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("AveragePerDuration %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *AveragePerDurationSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *AveragePerDurationSet) Get(name string) (*named.AveragePerDuration, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("AveragePerDuration %q not in set", name)
	}
	return ret, nil
}
