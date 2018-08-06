# registry
--
    import "github.com/KarelKubat/runtime-metrics/registry"

Package registry implements singletons of named sets.

The purpose of this is to uniquely register named metrics using singletons of
named sets. The metrics are then not only unique per named set, but also per
binary. This way, the reporter can pinpoint a metric by its name and be sure
that the metric is unique.

The typical usage is to create base metrics, register them using a unique name,
and start a reporting server. This is also shown in demo/demosrc/server.go:

    myAverage = base.NewAverage()
    err := registry.AddAverage("my_average", myAverage)
    if err != nil { ... } // name collision

    mySum := base.NewSum()
    err := registry.AddSum("my_sum", mySum)
    if err != nil { ... } // name collision

    // Now use myAverage and mySum to track whatever they should track.

There are methods to retrieve metrics by their name, which return an error when
the metric isn't in the registry. This is used by the reporter:

    avg, err := registry.GetAverage("my_average")
    if err != nil { ... } // not found

    sum, err := registry.GetSum("my_sum")
    if err != nil { ... } // not found

    // Now the values that are reported by avg.Report() and sum.Report()
    // can be published.

A sorted list of the names of all registered averages is retrieved using
AverageNames(), of all counts using CountNames(), etc.:

    for n := range registry.AverageNames() {
      fmt.Printf("An Average metric has the name %s\n", n)
      av, err := registry.GetAverage(n)
      if err != nil { ... } // internal error, should not happen
      val, n, until := av.Report()
      fmt.Printf("This average reports: average value %v, computed over %v values, until time %v\n",
        val, n, until)
    }

## Usage

#### func  AddAverage

```go
func AddAverage(name string, a *base.Average) *rtmerror.Error
```
AddAverage adds a reference to a base.Average to the registry, or returns an
error when a name collision occurs with an other metric.

#### func  AddAveragePerDuration

```go
func AddAveragePerDuration(name string, a *base.AveragePerDuration) *rtmerror.Error
```
AddAveragePerDuration adds a reference to a base.AveragePerDuration to the
registry, or returns an error when a name collision occurs with an other metric.

#### func  AddCount

```go
func AddCount(name string, a *base.Count) *rtmerror.Error
```
AddCount adds a reference to a base.Count to the registry, or returns an error
when a name collision occurs with an other metric.

#### func  AddCountPerDuration

```go
func AddCountPerDuration(name string, a *base.CountPerDuration) *rtmerror.Error
```
AddCountPerDuration adds a reference to a base.CountPerDuration to the registry,
or returns an error when a name collision occurs with an other metric.

#### func  AddSum

```go
func AddSum(name string, a *base.Sum) *rtmerror.Error
```
AddSum adds a reference to a base.Sum to the registry, or returns an error when
a name collision occurs with an other metric.

#### func  AddSumPerDuration

```go
func AddSumPerDuration(name string, a *base.SumPerDuration) *rtmerror.Error
```
AddSumPerDuration adds a reference to a base.SumPerDuration to the registry, or
returns an error when a name collision occurs with an other metric.

#### func  AverageBy

```go
func AverageBy(name string) (*base.Average, *rtmerror.Error)
```
AverageBy returns a reference to a registered base.Average, or a non-nil error
when the metric wasn't registered. The argument is the name to lookup.

#### func  AverageNames

```go
func AverageNames() []string
```
AverageNames returns a list of sorted strings of the names of the Averages in
the registry.

#### func  AveragePerDurationBy

```go
func AveragePerDurationBy(name string) (*base.AveragePerDuration, *rtmerror.Error)
```
AveragePerDurationBy returns a reference to a registered
base.AveragePerDuration, or a non-nil error when the metric wasn't registered.
The argument is the name to lookup.

#### func  AveragePerDurationNames

```go
func AveragePerDurationNames() []string
```
AveragePerDurationNames returns a list of sorted strings of all the names of
AveragePerDurations in the registry.

#### func  CountBy

```go
func CountBy(name string) (*base.Count, *rtmerror.Error)
```
CountBy returns a reference to a registered base.Count, or a non-nil error when
the metric wasn't registered. The argument is the name to lookup.

#### func  CountNames

```go
func CountNames() []string
```
CountNames returns a list of sorted strings of all the names of Counts in the
registry.

#### func  CountPerDurationBy

```go
func CountPerDurationBy(name string) (*base.CountPerDuration, *rtmerror.Error)
```
CountPerDurationBy returns a reference to a registered base.CountPerDuration, or
a non-nil error when the metric wasn't registered. The argument is the name to
lookup.

#### func  CountPerDurationNames

```go
func CountPerDurationNames() []string
```
CountPerDurationNames returns a list of sorted strings of all the names of
CountPerDurations in the registry.

#### func  SumBy

```go
func SumBy(name string) (*base.Sum, *rtmerror.Error)
```
SumBy returns a reference to a registered base.Sum, or a non-nil error when the
metric wasn't registered. The argument is the name to lookup.

#### func  SumNames

```go
func SumNames() []string
```
SumNames returns a list of sorted strings of all the names of Sums in the
registry.

#### func  SumPerDurationBy

```go
func SumPerDurationBy(name string) (*base.SumPerDuration, *rtmerror.Error)
```
SumPerDurationBy returns a reference to a registered base.SumPerDuration, or a
non-nil error when the metric wasn't registered. The argument is the name to
lookup.

#### func  SumPerDurationNames

```go
func SumPerDurationNames() []string
```
SumPerDurationNames returns a list of sorted strings of all the names of
SumPerDurations in the registry.
