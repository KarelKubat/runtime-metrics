package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

// AverageSet maps names to base.Average metrics.
type AverageSet struct {
	set   map[string]*base.Average
	mutex *sync.Mutex
}

// NewAverageSet returns an initialized AverageSet.
func NewAverageSet() *AverageSet {
	return &AverageSet{
		set:   map[string]*base.Average{},
		mutex: &sync.Mutex{},
	}
}

// Add registers a base.Average metric in the set.
func (set *AverageSet) Add(name string, a *base.Average) *rtmerror.Error {
	if set == nil {
		return rtmerror.NewError("attempt to add an Average to a non-initialized AverageSet")
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("Average %q already in set", name)
	}
	set.set[name] = a
	return nil
}

// Names returns all names of this set.
func (set *AverageSet) Names() []string {
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

// By returns a base.Average, identified by its name, or a non-nil error.
func (set *AverageSet) By(name string) (*base.Average, *rtmerror.Error) {
	if set == nil {
		return nil, nil
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("Average %q not in set", name)
	}
	return ret, nil
}
