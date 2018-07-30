package baseset

import (
	"fmt"
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
)

type CountSet struct {
	set   map[string]*base.Count
	mutex *sync.Mutex
}

func NewCountSet() *CountSet {
	return &CountSet{
		set:   map[string]*base.Count{},
		mutex: &sync.Mutex{},
	}
}

func (set *CountSet) Add(name string, a *base.Count) error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return fmt.Errorf("Count %q already in set", name)
	}
	set.set[name] = a
	return nil
}

func (set *CountSet) Names() []string {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CountSet) Get(name string) (*base.Count, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("Count %q not in set", name)
	}
	return ret, nil
}
