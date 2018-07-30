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

#### func (*Average) Mark

```go
func (a *Average) Mark(val float64)
```

#### func (*Average) Report

```go
func (a *Average) Report() (float64, int64)
```

#### func (*Average) Reset

```go
func (a *Average) Reset()
```

#### type AveragePerDuration

```go
type AveragePerDuration struct {
}
```


#### func  NewAveragePerDuration

```go
func NewAveragePerDuration(d time.Duration) *AveragePerDuration
```

#### func (*AveragePerDuration) Mark

```go
func (a *AveragePerDuration) Mark(val float64)
```

#### func (*AveragePerDuration) Report

```go
func (a *AveragePerDuration) Report() (float64, int64, time.Time)
```

#### func (*AveragePerDuration) Reset

```go
func (a *AveragePerDuration) Reset()
```

#### type Count

```go
type Count struct {
}
```


#### func  NewCount

```go
func NewCount() *Count
```

#### func (*Count) Mark

```go
func (c *Count) Mark()
```

#### func (*Count) Report

```go
func (c *Count) Report() int64
```

#### func (*Count) Reset

```go
func (c *Count) Reset()
```

#### type CountPerDuration

```go
type CountPerDuration struct {
}
```


#### func  NewCountPerDuration

```go
func NewCountPerDuration(d time.Duration) *CountPerDuration
```

#### func (*CountPerDuration) Mark

```go
func (c *CountPerDuration) Mark()
```

#### func (*CountPerDuration) Report

```go
func (c *CountPerDuration) Report() (int64, time.Time)
```

#### func (*CountPerDuration) Reset

```go
func (c *CountPerDuration) Reset()
```

#### type Sum

```go
type Sum struct {
}
```


#### func  NewSum

```go
func NewSum() *Sum
```

#### func (*Sum) Mark

```go
func (s *Sum) Mark(val float64)
```

#### func (*Sum) Report

```go
func (s *Sum) Report() (float64, int64)
```

#### func (*Sum) Reset

```go
func (s *Sum) Reset()
```

#### type SumPerDuration

```go
type SumPerDuration struct {
}
```


#### func  NewSumPerDuration

```go
func NewSumPerDuration(d time.Duration) *SumPerDuration
```

#### func (*SumPerDuration) Mark

```go
func (s *SumPerDuration) Mark(val float64)
```

#### func (*SumPerDuration) Report

```go
func (s *SumPerDuration) Report() (float64, int64, time.Time)
```

#### func (*SumPerDuration) Reset

```go
func (s *SumPerDuration) Reset()
```
