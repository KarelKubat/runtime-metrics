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
func AddAveragePerDuration(a *named.AveragePerDuration) error {
	return averagePerDurationSet.Add(a)
}
func AddCount(a *named.Count) error {
	return counterSet.Add(a)
}
func AddCountPerDuration(a *named.CountPerDuration) error {
	return counterPerDurationSet.Add(a)
}
func AddSum(a *named.Sum) error {
	return sumSet.Add(a)
}
func AddSumPerDuration(a *named.SumPerDuration) error {
	return sumPerDurationSet.Add(a)
}

func GetAverage(name string) (*named.Average, error) {
	return averageSet.Get(name)
}
func GetAveragePerDuration(name string) (*named.AveragePerDuration, error) {
	return averagePerDurationSet.Get(name)
}
func GetCount(name string) (*named.Count, error) {
	return counterSet.Get(name)
}
func GetCountPerDuration(name string) (*named.CountPerDuration, error) {
	return counterPerDurationSet.Get(name)
}
func GetSum(name string) (*named.Sum, error) {
	return sumSet.Get(name)
}
func GetSumPerDuration(name string) (*named.SumPerDuration, error) {
	return sumPerDurationSet.Get(name)
}

func AverageNames() []string {
	return averageSet.Names()
}
func AveragePerDurationNames() []string {
	return averagePerDurationSet.Names()
}
func CountNames() []string {
	return counterSet.Names()
}
func CountPerDurationNames() []string {
	return counterPerDurationSet.Names()
}
func SumNames() []string {
	return sumSet.Names()
}
func SumPerDurationNames() []string {
	return sumPerDurationSet.Names()
}
