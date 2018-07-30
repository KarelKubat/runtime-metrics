package registry

import (
	"github.com/KarelKubat/runtime-metrics/named"
	"github.com/KarelKubat/runtime-metrics/namedset"
)

var averageSet *namedset.AverageSet
var averagePerDurationSet *namedset.AveragePerDurationSet
var counterSet *namedset.CountSet
var counterPerDurationSet *namedset.CountPerDurationSet
var sumSet *namedset.SumSet
var sumPerDurationSet *namedset.SumPerDurationSet

func init() {
	averageSet = namedset.NewAverageSet()
	averagePerDurationSet = namedset.NewAveragePerDurationSet()
	counterSet = namedset.NewCountSet()
	counterPerDurationSet = namedset.NewCountPerDurationSet()
	sumSet = namedset.NewSumSet()
	sumPerDurationSet = namedset.NewSumPerDurationSet()
}

// AddAverage adds a reference to a named Average to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddAverage(a *named.Average) error {
	return averageSet.Add(a)
}

// AddAveragePerDuration adds a reference to a named AveragePerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddAveragePerDuration(a *named.AveragePerDuration) error {
	return averagePerDurationSet.Add(a)
}

// AddCount adds a reference to a named Count to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddCount(a *named.Count) error {
	return counterSet.Add(a)
}

// AddCountPerDuration adds a reference to a named CountPerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddCountPerDuration(a *named.CountPerDuration) error {
	return counterPerDurationSet.Add(a)
}

// AddSum adds a reference to a named Sum to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddSum(a *named.Sum) error {
	return sumSet.Add(a)
}

// AddSumPerDuration adds a reference to a named SumPerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddSumPerDuration(a *named.SumPerDuration) error {
	return sumPerDurationSet.Add(a)
}

// GetAverage returns a reference to a registered named.Average, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetAverage(name string) (*named.Average, error) {
	return averageSet.Get(name)
}

// GetAveragePerDuration returns a reference to a registered named.AveragePerDuration, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetAveragePerDuration(name string) (*named.AveragePerDuration, error) {
	return averagePerDurationSet.Get(name)
}

// GetCount returns a reference to a registered named.Count, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetCount(name string) (*named.Count, error) {
	return counterSet.Get(name)
}

// GetCountPerDuration returns a reference to a registered named.CountPerDuration, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetCountPerDuration(name string) (*named.CountPerDuration, error) {
	return counterPerDurationSet.Get(name)
}

// GetSum returns a reference to a registered named.Sum, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetSum(name string) (*named.Sum, error) {
	return sumSet.Get(name)
}

// GetSumPerDuration returns a reference to a registered named.SumPerDuration, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func GetSumPerDuration(name string) (*named.SumPerDuration, error) {
	return sumPerDurationSet.Get(name)
}

// AverageNames returns a list of sorted strings of all named Averages in the registry.
func AverageNames() []string {
	return averageSet.Names()
}

// AveragePerDurationNames returns a list of sorted strings of all named AveragePerDurations in the registry.
func AveragePerDurationNames() []string {
	return averagePerDurationSet.Names()
}

// CountNames returns a list of sorted strings of all named Counts in the registry.
func CountNames() []string {
	return counterSet.Names()
}

// CountPerDurationNames returns a list of sorted strings of all named CountPerDurations in the registry.
func CountPerDurationNames() []string {
	return counterPerDurationSet.Names()
}

// SumNames returns a list of sorted strings of all named Sums in the registry.
func SumNames() []string {
	return sumSet.Names()
}

// SumPerDurationNames returns a list of sorted strings of all named SumPerDurations in the registry.
func SumPerDurationNames() []string {
	return sumPerDurationSet.Names()
}
