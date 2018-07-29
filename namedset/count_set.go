package namedset

import (
	"fmt"
	"sort"

	named "github.com/KarelKubat/runtime-metrics/named"
)

type CountSet struct {
	set map[string]*named.Count
}

func NewCountSet() *CountSet {
	return &CountSet{
		set: map[string]*named.Count{},
	}
}

func (set *CountSet) Add(a *named.Count) error {
	if _, ok := set.set[a.Name()]; ok {
		return fmt.Errorf("Count %q already in set", a.Name())
	}
	set.set[a.Name()] = a
	return nil
}

func (set *CountSet) Names() []string {
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CountSet) Get(name string) (*named.Count, error) {
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("Count %q not in set", name)
	}
	return ret, nil
}
