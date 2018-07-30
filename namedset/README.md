# baseset
--
    import "github.com/KarelKubat/runtime-metrics/namedset"

baseset implements sets to uniquely register and identify base metrics.

This package is not relevant for normal usage. It is employed by the registry.

## Usage

#### type AveragePerDurationSet

```go
type AveragePerDurationSet struct {
}
```


#### func  NewAveragePerDurationSet

```go
func NewAveragePerDurationSet() *AveragePerDurationSet
```

#### func (*AveragePerDurationSet) Add

```go
func (set *AveragePerDurationSet) Add(name string, a *base.AveragePerDuration) error
```

#### func (*AveragePerDurationSet) Get

```go
func (set *AveragePerDurationSet) Get(name string) (*base.AveragePerDuration, error)
```

#### func (*AveragePerDurationSet) Names

```go
func (set *AveragePerDurationSet) Names() []string
```

#### type AverageSet

```go
type AverageSet struct {
}
```


#### func  NewAverageSet

```go
func NewAverageSet() *AverageSet
```

#### func (*AverageSet) Add

```go
func (set *AverageSet) Add(name string, a *base.Average) error
```

#### func (*AverageSet) Get

```go
func (set *AverageSet) Get(name string) (*base.Average, error)
```

#### func (*AverageSet) Names

```go
func (set *AverageSet) Names() []string
```

#### type CountPerDurationSet

```go
type CountPerDurationSet struct {
}
```


#### func  NewCountPerDurationSet

```go
func NewCountPerDurationSet() *CountPerDurationSet
```

#### func (*CountPerDurationSet) Add

```go
func (set *CountPerDurationSet) Add(name string, a *base.CountPerDuration) error
```

#### func (*CountPerDurationSet) Get

```go
func (set *CountPerDurationSet) Get(name string) (*base.CountPerDuration, error)
```

#### func (*CountPerDurationSet) Names

```go
func (set *CountPerDurationSet) Names() []string
```

#### type CountSet

```go
type CountSet struct {
}
```


#### func  NewCountSet

```go
func NewCountSet() *CountSet
```

#### func (*CountSet) Add

```go
func (set *CountSet) Add(name string, a *base.Count) error
```

#### func (*CountSet) Get

```go
func (set *CountSet) Get(name string) (*base.Count, error)
```

#### func (*CountSet) Names

```go
func (set *CountSet) Names() []string
```

#### type SumPerDurationSet

```go
type SumPerDurationSet struct {
}
```


#### func  NewSumPerDurationSet

```go
func NewSumPerDurationSet() *SumPerDurationSet
```

#### func (*SumPerDurationSet) Add

```go
func (set *SumPerDurationSet) Add(name string, a *base.SumPerDuration) error
```

#### func (*SumPerDurationSet) Get

```go
func (set *SumPerDurationSet) Get(name string) (*base.SumPerDuration, error)
```

#### func (*SumPerDurationSet) Names

```go
func (set *SumPerDurationSet) Names() []string
```

#### type SumSet

```go
type SumSet struct {
}
```


#### func  NewSumSet

```go
func NewSumSet() *SumSet
```

#### func (*SumSet) Add

```go
func (set *SumSet) Add(name string, a *base.Sum) error
```

#### func (*SumSet) Get

```go
func (set *SumSet) Get(name string) (*base.Sum, error)
```

#### func (*SumSet) Names

```go
func (set *SumSet) Names() []string
```
