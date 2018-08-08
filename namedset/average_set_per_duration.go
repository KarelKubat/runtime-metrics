package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

// AveragePerDurationSet maps names to base.AveragePerDuration metrics.
type AveragePerDurationSet struct {
	set   map[string]*base.AveragePerDuration
	mutex *sync.Mutex
}

// NewAveragePerDurationSet returns an initialized AveragePerDurationSet.
func NewAveragePerDurationSet() *AveragePerDurationSet {
	return &AveragePerDurationSet{
		set:   map[string]*base.AveragePerDuration{},
		mutex: &sync.Mutex{},
	}
}

// Add registers a base.AveragePerDuration in the set.
func (set *AveragePerDurationSet) Add(name string, a *base.AveragePerDuration) *rtmerror.Error {
	if set == nil {
		return rtmerror.NewError("attempt to add an AveragePerDuration to a non-initialized AveragePerDurationSet")
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()

	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("AveragePerDuration %q already in set", name)
	}
	set.set[name] = a
	return nil
}

// Names returns all names of this set.
func (set *AveragePerDurationSet) Names() []string {
	if set == nil {
		return []string{}
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()

	names := []string{}
	for name := range set.set {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// By returns a base.AveragePerDuration, identified by its name, or a non-nil
// error.
func (set *AveragePerDurationSet) By(name string) (*base.AveragePerDuration, *rtmerror.Error) {
	if set == nil {
		return nil, nil
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()

	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("AveragePerDuration %q not in set", name)
	}
	return ret, nil
}
