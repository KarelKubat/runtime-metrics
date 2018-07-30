# registry
--
    import "github.com/KarelKubat/runtime-metrics/registry"

registry implements singleton named sets.

The purpose of this is to uniquely register named metrics using singletons of
named sets. The metrics are then not only unique per named set, but also per
binary. This way, the reporter can pinpoint a metric by its name and be sure
that the metric is unique.

The typical usage is to create base metrics, register them using a unique name,
and start a reporting server. This is also shown in demo/server.go:

    err := registry.AddAverage("my_average", base.NewAverage())
    if err != nil { ... } // name collision

    err := registry.AddSum("my_sum", base.NewSum())
    if err != nil { ... } // name collision

There are methods to retrieve metrics by their name, which return an error when
the metric isn't in the registry. This is used by the reporter:

    avg, err := registry.GetAverage("my_average")
    if err != nil { ... } // not found

    sum, err := registry.GetSum("my_sum")
    if err != nil { ... } // not found

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
func AddAverage(name string, a *base.Average) error
```
AddAverage adds a reference to a base.Average to the registry, or returns an
error when a name collision occurs with an other metric.

#### func  AddAveragePerDuration

```go
func AddAveragePerDuration(name string, a *base.AveragePerDuration) error
```
AddAveragePerDuration adds a reference to a base.AveragePerDuration to the
registry, or returns an error when a name collision occurs with an other metric.

#### func  AddCount

```go
func AddCount(name string, a *base.Count) error
```
AddCount adds a reference to a base.Count to the registry, or returns an error
when a name collision occurs with an other metric.

#### func  AddCountPerDuration

```go
func AddCountPerDuration(name string, a *base.CountPerDuration) error
```
AddCountPerDuration adds a reference to a base.CountPerDuration to the registry,
or returns an error when a name collision occurs with an other metric.

#### func  AddSum

```go
func AddSum(name string, a *base.Sum) error
```
AddSum adds a reference to a base.Sum to the registry, or returns an error when
a name collision occurs with an other metric.

#### func  AddSumPerDuration

```go
func AddSumPerDuration(name string, a *base.SumPerDuration) error
```
AddSumPerDuration adds a reference to a base.SumPerDuration to the registry, or
returns an error when a name collision occurs with an other metric.

#### func  AverageNames

```go
func AverageNames() []string
```
AverageNames returns a list of sorted strings of the names of the Averages in
the registry.

#### func  AveragePerDurationNames

```go
func AveragePerDurationNames() []string
```
AveragePerDurationNames returns a list of sorted strings of all the names of
AveragePerDurations in the registry.

#### func  CountNames

```go
func CountNames() []string
```
CountNames returns a list of sorted strings of all the names of Counts in the
registry.

#### func  CountPerDurationNames

```go
func CountPerDurationNames() []string
```
CountPerDurationNames returns a list of sorted strings of all the names of
CountPerDurations in the registry.

#### func  GetAverage

```go
func GetAverage(name string) (*base.Average, error)
```
GetAverage returns a reference to a registered base.Average, or a non-nil error
when the metric wasn't registered. The argument is the name to lookup.

#### func  GetAveragePerDuration

```go
func GetAveragePerDuration(name string) (*base.AveragePerDuration, error)
```
GetAveragePerDuration returns a reference to a registered
base.AveragePerDuration, or a non-nil error when the metric wasn't registered.
The argument is the name to lookup.

#### func  GetCount

```go
func GetCount(name string) (*base.Count, error)
```
GetCount returns a reference to a registered base.Count, or a non-nil error when
the metric wasn't registered. The argument is the name to lookup.

#### func  GetCountPerDuration

```go
func GetCountPerDuration(name string) (*base.CountPerDuration, error)
```
GetCountPerDuration returns a reference to a registered base.CountPerDuration,
or a non-nil error when the metric wasn't registered. The argument is the name
to lookup.

#### func  GetSum

```go
func GetSum(name string) (*base.Sum, error)
```
GetSum returns a reference to a registered base.Sum, or a non-nil error when the
metric wasn't registered. The argument is the name to lookup.

#### func  GetSumPerDuration

```go
func GetSumPerDuration(name string) (*base.SumPerDuration, error)
```
GetSumPerDuration returns a reference to a registered base.SumPerDuration, or a
non-nil error when the metric wasn't registered. The argument is the name to
lookup.

#### func  SumNames

```go
func SumNames() []string
```
SumNames returns a list of sorted strings of all the names of Sums in the
registry.

#### func  SumPerDurationNames

```go
func SumPerDurationNames() []string
```
SumPerDurationNames returns a list of sorted strings of all the names of
SumPerDurations in the registry.
