package registry

import (
	"github.com/KarelKubat/runtime-metrics/named"
	"github.com/KarelKubat/runtime-metrics/namedset"
)

var averageSet *namedset.AverageSet
var averagePerDurationSet *namedset.AveragePerDurationSet
var counterSet *namedset.CounterSet
var counterPerDurationSet *namedset.CounterPerDurationSet
var sumSet *namedset.SumSet
var sumPerDurationSet *namedset.SumPerDurationSet

func init() {
	averageSet = namedset.NewAverageSet()
	averagePerDurationSet = namedset.NewAveragePerDurationSet()
	counterSet = namedset.NewCounterSet()
	counterPerDurationSet = namedset.NewCounterPerDurationSet()
	sumSet = namedset.NewSumSet()
	sumPerDurationSet = namedset.NewSumPerDurationSet()
}

func AddAverage(a *named.Average) error {
	return averageSet.Add(a)
}
func AddAveragePerDuration(a *named.AveragePerDuration) error {
	return averagePerDurationSet.Add(a)
}
func AddCounter(a *named.Counter) error {
	return counterSet.Add(a)
}
func AddCounterPerDuration(a *named.CounterPerDuration) error {
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
func GetCounter(name string) (*named.Counter, error) {
	return counterSet.Get(name)
}
func GetCounterPerDuration(name string) (*named.CounterPerDuration, error) {
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
func CounterNames() []string {
	return counterSet.Names()
}
func CounterPerDurationNames() []string {
	return counterPerDurationSet.Names()
}
func SumNames() []string {
	return sumSet.Names()
}
func SumPerDurationNames() []string {
	return sumPerDurationSet.Names()
}
