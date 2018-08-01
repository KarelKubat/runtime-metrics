package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

type SumPerDurationSet struct {
	set   map[string]*base.SumPerDuration
	mutex *sync.Mutex
}

func NewSumPerDurationSet() *SumPerDurationSet {
	return &SumPerDurationSet{
		set:   map[string]*base.SumPerDuration{},
		mutex: &sync.Mutex{},
	}
}

func (set *SumPerDurationSet) Add(name string, a *base.SumPerDuration) error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("SumPerDuration %q already in set", name)
	}
	set.set[name] = a
	return nil
}

func (set *SumPerDurationSet) Names() []string {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (set *SumPerDurationSet) Get(name string) (*base.SumPerDuration, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("SumPerDuration %q not in set", name)
	}
	return ret, nil
}
