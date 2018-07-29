package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type CountPerDurationSet struct {
	set map[string]*named.CountPerDuration
}

func NewCountPerDurationSet() *CountPerDurationSet {
	return &CountPerDurationSet{
		set: map[string]*named.CountPerDuration{},
	}
}

func (set *CountPerDurationSet) Add(a *named.CountPerDuration) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("CountPerDuration %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *CountPerDurationSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CountPerDurationSet) Get(name string) (*named.CountPerDuration, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("CountPerDuration %q not in set", name)
	}
	return ret, nil
}
