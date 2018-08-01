# base
--
    import "github.com/KarelKubat/runtime-metrics/base"

Package base defines the most basic metric types.

The most basic types are:

    base.Count                // records ticks
    base.Sum                  // adds float64 values
    base.Average              // just a complimentary type very similar to Sum

Besides these, *PerDuration exist, to record over a given period, which gives
the additional:

    base.CountPerDuration     // count over a period
    base.SumPerDuration       // sum over a period
    base.AveragePerDuration   // average over a period

New metric instances are created using New*, such as NewCount(), NewSum(), and
so on. Marking an event in a metric is done using the Mark() method, which takes
different arguments, depending on the metric. Resetting a metric is done using
Reset(). Seeing what the state of the metric is, is done using Report(), which
returns different values, depeinding on the metric.

Example:

    av := base.NewAverage()
    av.Mark(1.0)
    av.Mark(2.0)
    average, n := av.Report()
    fmt.Printf("average is %v, over %v values\n", average, n) // 1.5 over 2 values

    av.Reset()
    av.Mark(4.0)
    av.Mark(5.0)
    average, n = av.Report()
    fmt.Printf("average is %v, over %v values\n", average, n) // 4.5 over 2 values

All *PerDuration metrics are created with a time.Duration specifier, which is
the "width" of the observation window. NOTE THAT when Report() is called on
these metrics, then values are returned that were obtained somewhere in the
past. This is best illustrated in a simple diagram:

           av := base.NewAveragePerDuration(time.Second)
    T=0s
     |     av.Mark(1.0)
     |     av.Mark(2.0)
     |     val0, n0, until0 := av.Report()   // state before T=0; val0 is 0.0
    T=1s
     |     av.Mark(20.0)
     |     av.Mark(10.0)
     |     val1, n1, until1 := av.Report()   // state between T=0 and T=1; val1 is 1.5
    T=2s
     |     av.Mark(200.0)
     |     av.Mark(100.0)
     |     val2, n2, until2 := av.Report()   // state between T=1 and T=2; val2 is 15.0
    T=3s
     |     val3, n3, until3 := av.Report     // state between T=2 and T=3; val3 is 150.0
     |

Here, val0 and n0 will both be zero, because the very first period is not
finished yet when Report() is called. Once one second has expired (T=1), val1
will report 1.5, and the marked values 10.0 and 20.0 are again not taken into
account, as they fall into the currently updating period. Only after T=2 has
passed, Report() will return 15.0 as the average.

For all *PerDuration metrics, Report() also returns the timestamp until which
the marked values are taken into account.

## Usage

#### type Average

```go
type Average struct {
}
```

Average is the metric type for averages.

#### func  NewAverage

```go
func NewAverage() *Average
```
NewAverage returns a reference to this metric type.

#### func (*Average) Mark

```go
func (a *Average) Mark(val float64)
```
Mark marks the occurrence of a floating point value.

#### func (*Average) Report

```go
func (a *Average) Report() (float64, int64)
```
Report returns the average and number of observed values.

#### func (*Average) Reset

```go
func (a *Average) Reset()
```
Reset resets the metric.

#### type AveragePerDuration

```go
type AveragePerDuration struct {
}
```

AveragePerDuration is the metric type for averages over a time span.

#### func  NewAveragePerDuration

```go
func NewAveragePerDuration(d time.Duration) *AveragePerDuration
```
NewAveragePerDuration returns a reference to this metric type.

#### func (*AveragePerDuration) Mark

```go
func (a *AveragePerDuration) Mark(val float64)
```
Mark marks an observation of a floating point value.

#### func (*AveragePerDuration) Report

```go
func (a *AveragePerDuration) Report() (float64, int64, time.Time)
```
Report returns the average, number of observed values, and time until which the
avarage was computed. The observation started at the returned timestamp minus
the duration.

#### func (*AveragePerDuration) Reset

```go
func (a *AveragePerDuration) Reset()
```
Reset resets the metric.

#### type Count

```go
type Count struct {
}
```

Count is the metric type for counters.

#### func  NewCount

```go
func NewCount() *Count
```
NewCount returns a reference to this metric type.

#### func (*Count) Mark

```go
func (c *Count) Mark()
```
Mark marks the occurence of a "tick".

#### func (*Count) Report

```go
func (c *Count) Report() int64
```
Report returns the number of observed "ticks".

#### func (*Count) Reset

```go
func (c *Count) Reset()
```
Reset resets the metric.

#### type CountPerDuration

```go
type CountPerDuration struct {
}
```

CountPerDuration is the metric type for counts over a time span.

#### func  NewCountPerDuration

```go
func NewCountPerDuration(d time.Duration) *CountPerDuration
```
NewCountPerDuration returns a reference to this metric.

#### func (*CountPerDuration) Mark

```go
func (c *CountPerDuration) Mark()
```
Mark marks the occurrence of a "tick".

#### func (*CountPerDuration) Report

```go
func (c *CountPerDuration) Report() (int64, time.Time)
```
Report returns the number of observed "ticks" and the time until which the count
was maintained.

#### func (*CountPerDuration) Reset

```go
func (c *CountPerDuration) Reset()
```
Reset resets the metric.

#### type Sum

```go
type Sum struct {
}
```

Sum is the metric type for sums.

#### func  NewSum

```go
func NewSum() *Sum
```
NewSum returns a reference to this metric type.

#### func (*Sum) Mark

```go
func (s *Sum) Mark(val float64)
```
Mark adds the occurrence of a floating point value.

#### func (*Sum) Report

```go
func (s *Sum) Report() (float64, int64)
```
Report returns the sum and number of observed values.

#### func (*Sum) Reset

```go
func (s *Sum) Reset()
```
Reset resets the metric.

#### type SumPerDuration

```go
type SumPerDuration struct {
}
```

SumPerDuration is the metric type for sums over a time span.

#### func  NewSumPerDuration

```go
func NewSumPerDuration(d time.Duration) *SumPerDuration
```
NewSumPerDuration returns a reference to this metric type.

#### func (*SumPerDuration) Mark

```go
func (s *SumPerDuration) Mark(val float64)
```
Mark marks the occurrence of a floating point value.

#### func (*SumPerDuration) Report

```go
func (s *SumPerDuration) Report() (float64, int64, time.Time)
```
Report returns the sum of the observed values, the number of observed values,
and the time until which the sum was computed.

#### func (*SumPerDuration) Reset

```go
func (s *SumPerDuration) Reset()
```
Reset resets this metric.
