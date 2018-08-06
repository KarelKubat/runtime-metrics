package namedset

import (
	"sort"
	"sync"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
)

// CountSet maps names to base.Count metrics.
type CountSet struct {
	set   map[string]*base.Count
	mutex *sync.Mutex
}

// NewCountSet returns an initialized CountSet.
func NewCountSet() *CountSet {
	return &CountSet{
		set:   map[string]*base.Count{},
		mutex: &sync.Mutex{},
	}
}

// Add registers a base.Count metric in the set.
func (set *CountSet) Add(name string, a *base.Count) *rtmerror.Error {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	if _, ok := set.set[name]; ok {
		return rtmerror.NewError("Count %q already in set", name)
	}
	set.set[name] = a
	return nil
}

// Names returns all names of this set.
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

// By returns a base.Count, identified by its name, or a non-nil error.
func (set *CountSet) By(name string) (*base.Count, *rtmerror.Error) {
	set.mutex.Lock()
	defer set.mutex.Unlock()
	ret, ok := set.set[name]
	if !ok {
		return nil, rtmerror.NewError("Count %q not in set", name)
	}
	return ret, nil
}
