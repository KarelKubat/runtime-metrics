package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

type AverageSet struct {
	set   map[string]*base.Average
	mutex *sync.Mutex
}

func NewAverageSet() *AverageSet {
	return &AverageSet{
		set:   map[string]*base.Average{},
		mutex: &sync.Mutex{},
	}
}

func (set *AverageSet) Add(name string, a *base.Average) error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("Average %q already in set", name)
	}
	set.set[name] = a
	return nil
}

func (set *AverageSet) Names() []string {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *AverageSet) Get(name string) (*base.Average, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("Average %q not in set", name)
	}
	return ret, nil
}
