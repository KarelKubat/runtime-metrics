# namedset
--
    import "github.com/KarelKubat/runtime-metrics/namedset"

Package namedset implements sets to uniquely register and identify base metrics.

This package is not relevant for normal usage. It is employed by the registry.

## Usage

#### type AveragePerDurationSet

```go
type AveragePerDurationSet struct {
}
```

AveragePerDurationSet maps names to base.AveragePerDuration metrics.

#### func  NewAveragePerDurationSet

```go
func NewAveragePerDurationSet() *AveragePerDurationSet
```
NewAveragePerDurationSet returns an initialized AveragePerDurationSet.

#### func (*AveragePerDurationSet) Add

```go
func (set *AveragePerDurationSet) Add(name string, a *base.AveragePerDuration) *rtmerror.Error
```
Add registers a base.AveragePerDuration in the set.

#### func (*AveragePerDurationSet) By

```go
func (set *AveragePerDurationSet) By(name string) (*base.AveragePerDuration, *rtmerror.Error)
```
By returns a base.AveragePerDuration, identified by its name, or a non-nil
error.

#### func (*AveragePerDurationSet) Names

```go
func (set *AveragePerDurationSet) Names() []string
```
Names returns all names of this set.

#### type AverageSet

```go
type AverageSet struct {
}
```

AverageSet maps names to base.Average metrics.

#### func  NewAverageSet

```go
func NewAverageSet() *AverageSet
```
NewAverageSet returns an initialized AverageSet.

#### func (*AverageSet) Add

```go
func (set *AverageSet) Add(name string, a *base.Average) *rtmerror.Error
```
Add registers a base.Average metric in the set.

#### func (*AverageSet) By

```go
func (set *AverageSet) By(name string) (*base.Average, *rtmerror.Error)
```
By returns a base.Average, identified by its name, or a non-nil error.

#### func (*AverageSet) Names

```go
func (set *AverageSet) Names() []string
```
Names returns all names of this set.

#### type CountPerDurationSet

```go
type CountPerDurationSet struct {
}
```

CountPerDurationSet maps names to base.CountPerDuration metrics.

#### func  NewCountPerDurationSet

```go
func NewCountPerDurationSet() *CountPerDurationSet
```
NewCountPerDurationSet returns an initialized CountPerDurationSet.

#### func (*CountPerDurationSet) Add

```go
func (set *CountPerDurationSet) Add(name string, a *base.CountPerDuration) *rtmerror.Error
```
Add registers a base.CountPerDuration metric in the set.

#### func (*CountPerDurationSet) By

```go
func (set *CountPerDurationSet) By(name string) (*base.CountPerDuration, *rtmerror.Error)
```
By returns a base.CountPerDuration, identified by its name, or a non-nil error.

#### func (*CountPerDurationSet) Names

```go
func (set *CountPerDurationSet) Names() []string
```
Names returns all names in this set.

#### type CountSet

```go
type CountSet struct {
}
```

CountSet maps names to base.Count metrics.

#### func  NewCountSet

```go
func NewCountSet() *CountSet
```
NewCountSet returns an initialized CountSet.

#### func (*CountSet) Add

```go
func (set *CountSet) Add(name string, a *base.Count) *rtmerror.Error
```
Add registers a base.Count metric in the set.

#### func (*CountSet) By

```go
func (set *CountSet) By(name string) (*base.Count, *rtmerror.Error)
```
By returns a base.Count, identified by its name, or a non-nil error.

#### func (*CountSet) Names

```go
func (set *CountSet) Names() []string
```
Names returns all names of this set.

#### type SumPerDurationSet

```go
type SumPerDurationSet struct {
}
```

SumPerDurationSet maps names to base.SumPerDuration metrics.

#### func  NewSumPerDurationSet

```go
func NewSumPerDurationSet() *SumPerDurationSet
```
NewSumPerDurationSet returns an initialized SumPerDurationSet.

#### func (*SumPerDurationSet) Add

```go
func (set *SumPerDurationSet) Add(name string, a *base.SumPerDuration) *rtmerror.Error
```
Add registers a base.SumPerDuration metric in the set.

#### func (*SumPerDurationSet) By

```go
func (set *SumPerDurationSet) By(name string) (*base.SumPerDuration, *rtmerror.Error)
```
By returns a base.SumPerDuration, identified by its name, or a non-nil error.

#### func (*SumPerDurationSet) Names

```go
func (set *SumPerDurationSet) Names() []string
```
Names returns all names of this set.

#### type SumSet

```go
type SumSet struct {
}
```

SumSet maps names to base.Sum metrics.

#### func  NewSumSet

```go
func NewSumSet() *SumSet
```
NewSumSet returns an initialized SumSet.

#### func (*SumSet) Add

```go
func (set *SumSet) Add(name string, a *base.Sum) *rtmerror.Error
```
Add registers a base.Sum metric in the set.

#### func (*SumSet) By

```go
func (set *SumSet) By(name string) (*base.Sum, *rtmerror.Error)
```
By returns a base.Sum, identified by its name, or a non-nil error.

#### func (*SumSet) Names

```go
func (set *SumSet) Names() []string
```
Names returns all names of this set.
