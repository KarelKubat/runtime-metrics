package registry

import (
	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/namedset"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
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

// AddAverage adds a reference to a base.Average to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddAverage(name string, a *base.Average) *rtmerror.Error {
	return averageSet.Add(name, a)
}

// AddAveragePerDuration adds a reference to a base.AveragePerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddAveragePerDuration(name string, a *base.AveragePerDuration) *rtmerror.Error {
	return averagePerDurationSet.Add(name, a)
}

// AddCount adds a reference to a base.Count to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddCount(name string, a *base.Count) *rtmerror.Error {
	return counterSet.Add(name, a)
}

// AddCountPerDuration adds a reference to a base.CountPerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddCountPerDuration(name string, a *base.CountPerDuration) *rtmerror.Error {
	return counterPerDurationSet.Add(name, a)
}

// AddSum adds a reference to a base.Sum to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddSum(name string, a *base.Sum) *rtmerror.Error {
	return sumSet.Add(name, a)
}

// AddSumPerDuration adds a reference to a base.SumPerDuration to the registry,
// or returns an error when a name collision occurs with an other
// metric.
func AddSumPerDuration(name string, a *base.SumPerDuration) *rtmerror.Error {
	return sumPerDurationSet.Add(name, a)
}

// AverageBy returns a reference to a registered base.Average, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func AverageBy(name string) (*base.Average, *rtmerror.Error) {
	return averageSet.By(name)
}

// AveragePerDurationBy returns a reference to a registered base.AveragePerDuration,
// or a non-nil error when the metric wasn't registered. The argument is the name to lookup.
func AveragePerDurationBy(name string) (*base.AveragePerDuration, *rtmerror.Error) {
	return averagePerDurationSet.By(name)
}

// CountBy returns a reference to a registered base.Count, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func CountBy(name string) (*base.Count, *rtmerror.Error) {
	return counterSet.By(name)
}

// CountPerDurationBy returns a reference to a registered base.CountPerDuration, or a non-nil
// error when the metric wasn't registered. The argument is the name to lookup.
func CountPerDurationBy(name string) (*base.CountPerDuration, *rtmerror.Error) {
	return counterPerDurationSet.By(name)
}

// SumBy returns a reference to a registered base.Sum, or a non-nil error when
// the metric wasn't registered. The argument is the name to lookup.
func SumBy(name string) (*base.Sum, *rtmerror.Error) {
	return sumSet.By(name)
}

// SumPerDurationBy returns a reference to a registered base.SumPerDuration, or a non-nil error
// when the metric wasn't registered. The argument is the name to lookup.
func SumPerDurationBy(name string) (*base.SumPerDuration, *rtmerror.Error) {
	return sumPerDurationSet.By(name)
}

// AverageNames returns a list of sorted strings of the names of the Averages in the registry.
func AverageNames() []string {
	return averageSet.Names()
}

// AveragePerDurationNames returns a list of sorted strings of all the names of AveragePerDurations in the registry.
func AveragePerDurationNames() []string {
	return averagePerDurationSet.Names()
}

// CountNames returns a list of sorted strings of all the names of Counts in the registry.
func CountNames() []string {
	return counterSet.Names()
}

// CountPerDurationNames returns a list of sorted strings of all the names of CountPerDurations in the registry.
func CountPerDurationNames() []string {
	return counterPerDurationSet.Names()
}

// SumNames returns a list of sorted strings of all the names of Sums in the registry.
func SumNames() []string {
	return sumSet.Names()
}

// SumPerDurationNames returns a list of sorted strings of all the names of SumPerDurations in the registry.
func SumPerDurationNames() []string {
	return sumPerDurationSet.Names()
}
