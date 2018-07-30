package baseset

import (
	"fmt"
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
)

type CountPerDurationSet struct {
	set   map[string]*base.CountPerDuration
	mutex *sync.Mutex
}

func NewCountPerDurationSet() *CountPerDurationSet {
	return &CountPerDurationSet{
		set:   map[string]*base.CountPerDuration{},
		mutex: &sync.Mutex{},
	}
}

func (set *CountPerDurationSet) Add(name string, a *base.CountPerDuration) error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return fmt.Errorf("CountPerDuration %q already in set", name)
	}
	set.set[name] = a
	return nil
}

func (set *CountPerDurationSet) Names() []string {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *CountPerDurationSet) Get(name string) (*base.CountPerDuration, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, fmt.Errorf("CountPerDuration %q not in set", name)
	}
	return ret, nil
}
