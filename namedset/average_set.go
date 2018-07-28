package namedset

import (
	"fmt"
	"sort"

	"github.com/KarelKubat/runtime-metrics/named"
)

type AverageSet struct {
	set map[string]*named.Average
}

func NewAverageSet() *AverageSet {
	return &AverageSet{
		set: map[string]*named.Average{},
	}
}

func (set *AverageSet) Add(a *named.Average) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("Average %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *AverageSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *AverageSet) Get(name string) (*named.Average, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("Average %q not in set", name)
	}
	return ret, nil
}
