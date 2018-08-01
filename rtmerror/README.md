# rtmerror
--
    import "github.com/KarelKubat/runtime-metrics/rtmerror"

Package rtmerror implements the error type for runtime-metrics.

Errors in runtime-metrics are a custom error type, rtmerror.Error. You will
normally not care, except when checking whether an error is retryable. This only
occurs during network calls, and normally the backoff policy of the reporting
client takes care of retries anyway.

## Usage

#### type Error

```go
type Error struct {
}
```

Error is the error type in runtime-metrics.

#### func  NewError

```go
func NewError(f string, a ...interface{}) *Error
```
NewError initializes an Error return struct with a description, and marks the
error as non-retryable. Example:

    err := rtmerror.NewError("%d is the answer", 42)

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error implements the error interface and returns the error's description, with
underlying errors concatenated and between (). Example:

    err1 := fmt.Errorf("hello %s", world)
    err2 := fmt.Errorf("pi is not %.2f", 3.14)
    err := rtmerror.NewError("e is not %.2f", 2.71").
      WithError(err1).WithError(err2)
    fmt.Printf("%v\n", err)
    // output:
    // e is not 2.71 (hello world) (pi is not 3.14)

#### func (*Error) Retryable

```go
func (e *Error) Retryable() bool
```
Retryable returns true when the error is retryable according to its argument.

#### func (*Error) WithError

```go
func (e *Error) WithError(err error) *Error
```
WithError adds an underlying error to the Error structure. Example:

    var err1, err2 error
    err := rtmerror.NewError("something bad happened").
      WithError(err1).WithError(err2)

#### func (*Error) WithRetryable

```go
func (e *Error) WithRetryable(r bool) *Error
```
WithRetryable marks the error as retryable.

    err := rtmerror.NewError("network error").WithRetryable(true)
