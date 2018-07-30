# named
--
    import "github.com/KarelKubat/runtime-metrics/named"

named wraps the thread-safe metrics and assigns a name to them.

The constructors all take a string argument, and return a named metric:

    avg := named.NewAverage("my-average")
    avg_per_s := named.NewAveragePerDuration("my-average-per-sec", Time.Second)

    cnt := named.NewCount("my-count")
    cnt_per_s := named.NewcountPerDuration("my-count-per-sec", Time.Second)

    sum := named.NewSum("my-sum")
    sum_per_s := named.NewSumPerDuration("my-sum-per-sec", Time.Second)

The name is returned using method Name():

    fmt.Printf("The metric for averages is named %s\n", avg.Name())

The remainder of the API is identical to the underlying thread-safe metric
(which in its turn is identical to the underlying base metric); there are
methods Mark(), Report() and Reset(). Refer to the base package for a
description.

## Usage

#### type Average

```go
type Average struct {
}
```


#### func  NewAverage

```go
func NewAverage(n string) *Average
```

#### func (*Average) Mark

```go
func (a *Average) Mark(val float64)
```

#### func (*Average) Name

```go
func (a *Average) Name() string
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
func NewAveragePerDuration(n string, d time.Duration) *AveragePerDuration
```

#### func (*AveragePerDuration) Mark

```go
func (a *AveragePerDuration) Mark(val float64)
```

#### func (*AveragePerDuration) Name

```go
func (a *AveragePerDuration) Name() string
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
func NewCount(n string) *Count
```

#### func (*Count) Mark

```go
func (c *Count) Mark()
```

#### func (*Count) Name

```go
func (c *Count) Name() string
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
func NewCountPerDuration(n string, d time.Duration) *CountPerDuration
```

#### func (*CountPerDuration) Mark

```go
func (c *CountPerDuration) Mark()
```

#### func (*CountPerDuration) Name

```go
func (c *CountPerDuration) Name() string
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
func NewSum(n string) *Sum
```

#### func (*Sum) Mark

```go
func (c *Sum) Mark(val float64)
```

#### func (*Sum) Name

```go
func (c *Sum) Name() string
```

#### func (*Sum) Report

```go
func (c *Sum) Report() (float64, int64)
```

#### func (*Sum) Reset

```go
func (c *Sum) Reset()
```

#### type SumPerDuration

```go
type SumPerDuration struct {
}
```


#### func  NewSumPerDuration

```go
func NewSumPerDuration(n string, d time.Duration) *SumPerDuration
```

#### func (*SumPerDuration) Mark

```go
func (c *SumPerDuration) Mark(val float64)
```

#### func (*SumPerDuration) Name

```go
func (c *SumPerDuration) Name() string
```

#### func (*SumPerDuration) Report

```go
func (c *SumPerDuration) Report() (float64, int64, time.Time)
```

#### func (*SumPerDuration) Reset

```go
func (c *SumPerDuration) Reset()
```
