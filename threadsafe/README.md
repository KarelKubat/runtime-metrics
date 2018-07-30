# threadsafe
--
    import "github.com/KarelKubat/runtime-metrics/threadsafe"


## Usage

#### type Average

```go
type Average struct {
}
```


#### func  NewAverage

```go
func NewAverage() *Average
```
NewAverage returns a reference to an initialized threadsafe.Average.

#### func (*Average) Mark

```go
func (a *Average) Mark(val float64)
```
Mark registers a float64 observation.

#### func (*Average) Report

```go
func (a *Average) Report() (float64, int64)
```
Report returns the average over all observations and the number of cases.

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


#### func  NewAveragePerDuration

```go
func NewAveragePerDuration(d time.Duration) *AveragePerDuration
```
NewAveragePerDuration returns a reference to an initialized
threadsafe.AveragePerDuration. The argument is the time window, such as
10*time.Second.

#### func (*AveragePerDuration) Mark

```go
func (a *AveragePerDuration) Mark(val float64)
```
Mark registers a float64 observation.

#### func (*AveragePerDuration) Report

```go
func (a *AveragePerDuration) Report() (float64, int64, time.Time)
```
Report returns the average over all observations, the number of cases, and the
ending period for this average.

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


#### func  NewCount

```go
func NewCount() *Count
```
Newcount returns a reference to an initialized threadsafe.Count.

#### func (*Count) Mark

```go
func (c *Count) Mark()
```
Mark registers an observation (a "tick") and internally increments the count.

#### func (*Count) Report

```go
func (c *Count) Report() int64
```
Report returns the number of observations.

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


#### func  NewCountPerDuration

```go
func NewCountPerDuration(d time.Duration) *CountPerDuration
```
NewCountPerDuration returns a reference to an initialized
threadsafe.CountPerDuration. The argument is the time window over which counting
will apply.

#### func (*CountPerDuration) Mark

```go
func (c *CountPerDuration) Mark()
```
Mark registers an observation.

#### func (*CountPerDuration) Report

```go
func (c *CountPerDuration) Report() (int64, time.Time)
```
Report returns the number of observations and the ending time of counting.

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


#### func  NewSum

```go
func NewSum() *Sum
```
NewSum returns a reference to an initialized threadsafe.Sum.

#### func (*Sum) Mark

```go
func (s *Sum) Mark(val float64)
```
Mark registers a float64 observation.

#### func (*Sum) Report

```go
func (s *Sum) Report() (float64, int64)
```
Report returns the sum of observations and the number of cases.

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


#### func  NewSumPerDuration

```go
func NewSumPerDuration(d time.Duration) *SumPerDuration
```
NewSumPerDuration returns a reference to an initialized
threadsafe.SumPerDuration. The argument is the time window for summing.

#### func (*SumPerDuration) Mark

```go
func (s *SumPerDuration) Mark(val float64)
```
Mark registers a float64 observation.

#### func (*SumPerDuration) Report

```go
func (s *SumPerDuration) Report() (float64, int64, time.Time)
```
Report returns the sum, number of cases and end time.

#### func (*SumPerDuration) Reset

```go
func (s *SumPerDuration) Reset()
```
Reset resets the metric.
