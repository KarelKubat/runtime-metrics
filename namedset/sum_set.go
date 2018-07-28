package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type SumSet struct {
	set map[string]*named.Sum
}

func NewSumSet() *SumSet {
	return &SumSet{
		set: map[string]*named.Sum{},
	}
}

func (set *SumSet) Add(a *named.Sum) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("Sum %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *SumSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *SumSet) Get(name string) (*named.Sum, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("Sum %q not in set", name)
	}
	return ret, nil
}
