package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

// SumSet maps names to base.Sum metrics.
type SumSet struct {
	set   map[string]*base.Sum
	mutex *sync.Mutex
}

// NewSumSet returns an initialized SumSet.
func NewSumSet() *SumSet {
	return &SumSet{
		set:   map[string]*base.Sum{},
		mutex: &sync.Mutex{},
	}
}

// Add registers a base.Sum metric in the set.
func (set *SumSet) Add(name string, a *base.Sum) *rtmerror.Error {
	if set == nil {
		return rtmerror.NewError("attempt to Add a Sum to a non-initialized SumSet")
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("Sum %q already in set", name)
	}
	set.set[name] = a
	return nil
}

// Names returns all names of this set.
func (set *SumSet) Names() []string {
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

// By returns a base.Sum, identified by its name, or a non-nil error.
func (set *SumSet) By(name string) (*base.Sum, *rtmerror.Error) {
	if set == nil {
		return nil, nil
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("Sum %q not in set", name)
	}
	return ret, nil
}
