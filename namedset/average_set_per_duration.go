package namedset

import (
	"fmt"
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
)

type AveragePerDurationSet struct {
	set   map[string]*base.AveragePerDuration
	mutex *sync.Mutex
}

func NewAveragePerDurationSet() *AveragePerDurationSet {
	return &AveragePerDurationSet{
		set:   map[string]*base.AveragePerDuration{},
		mutex: &sync.Mutex{},
	}
}

func (set *AveragePerDurationSet) Add(name string, a *base.AveragePerDuration) error {
	if _, ok := set.set[name]; ok {
		return fmt.Errorf("AveragePerDuration %q already in set", name)
	}
	set.set[name] = a
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

func (set *AveragePerDurationSet) Get(name string) (*base.AveragePerDuration, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("AveragePerDuration %q not in set", name)
	}
	return ret, nil
}
