package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

// CountPerDurationSet maps names to base.CountPerDuration metrics.
type CountPerDurationSet struct {
	set   map[string]*base.CountPerDuration
	mutex *sync.Mutex
}

// NewCountPerDurationSet returns an initialized CountPerDurationSet.
func NewCountPerDurationSet() *CountPerDurationSet {
	return &CountPerDurationSet{
		set:   map[string]*base.CountPerDuration{},
		mutex: &sync.Mutex{},
	}
}

// Add registers a base.CountPerDuration metric in the set.
func (set *CountPerDurationSet) Add(name string, a *base.CountPerDuration) error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("CountPerDuration %q already in set", name)
	}
	set.set[name] = a
	return nil
}

// Names returns all names in this set.
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

// By returns a base.CountPerDuration, identified by its name, or a non-nil error.
func (set *CountPerDurationSet) By(name string) (*base.CountPerDuration, error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("CountPerDuration %q not in set", name)
	}
	return ret, nil
}
