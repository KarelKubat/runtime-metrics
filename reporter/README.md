# reporter
--
    import "github.com/KarelKubat/runtime-metrics/reporter"

reporter implements a metrics reporting server and client.


### The Server

The server is started using reporter.StartReporter(addr), where addr defines the
IP and port to listen, separated by :. The IP can be left out. Example:

    err := reporter.StartReporter(":9000")    // bind to TCP port 9000
    if err != nil { ... }                     // probably port already taken

Typically this will be wrapped in a go function, and colocated with
instantiating metrics:

    func checkErr(err) {
      if err != nil {
        handleError()
      }
    }

    ...

    checkErr(registry.AddAverage(named.NewAverage("average-metric")))
    checkErr(registry.AddSumPerDuration(named.NewSumPerDuration("sum-per-minute", 30 * time.Minute)))

    go func() {
      checkErr(reporter.StartReporter(":9000")
    }()


### The Client

The client is instantiated using reporter.NewClient(addr), where again addr is a
string defining the IP address and TCP port:

    c, err := reporter.NewClient(":9000")
    checkErr(err)
    defer c.Close()

The client has a number of handy methods that contact the client to discover the
names of metrics or to fetch metric values:

    allNames, err := c.AllNames()
    // allNames.Averages is an array of strings (names) of all Average-type metrics
    // allNames.AveragesByDuration is an array of strings (names) of all AveragePerDuration-type metrics
    // allNames.Counts is an array of strings (names) of all Count-type metrics
    // allNames.CountsPerDuration is an array of strings (names) of all CountPerDuration-type metrics
    // allNames.Sums is an array of strings (names) of all Sum-type metrics
    // allNames.SumsPerDuration is an array of strings (names) of all SumPerDuration-type metrics

In order to fetch the values of a metric, the client calls c.Average(name),
c.Sum(name) etc. The returned values are always what the base type returns, and
an error:

    avg, n, err := c.Average("my-average")
    // val is the average
    // n   is the number of observations
    // err is nil or an error

See demo/client.go for an example.

## Usage

#### func  StartReporter

```go
func StartReporter(addr string) error
```
StartReporter starts the reporting server or returns a non-nil error. The
argument is an address in the format "ip:port", where "ip" is optional. This is
the binding address.

#### type AllNamesReturn

```go
type AllNamesReturn struct {
	Averages            []string // names of averages
	AveragesPerDuration []string // names of averages per duration
	Counts              []string // names of counts
	CountsPerDuration   []string // names of counts per duration
	Sums                []string // names of sums
	SumsPerDuration     []string // names of sums per duration
}
```

AllNamesReturn is returned by AllNames and lists the names of averages, averages
per duration, etc.

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(addr string) (*Client, error)
```
NewClient returns an initialized client, or a non-nil error. The addr argument
is the "ip:port" to connect to; "ip" being optional (defaults to localhost).

#### func (*Client) AllNames

```go
func (c *Client) AllNames() (*AllNamesReturn, error)
```
AllNames returns a list of server-registered named metrics, or a non-nil error.

Example:

    overview, err := client.AllNameS()
    if err != nil { ... }
    for _, n := range overview.Averages {
      fmt.Printf("There is a named average metric called %s\n", n)
    }
    // and so on with overview.AveragesPerDuration, overview.Counts,
    // overview.CountsPerDuration, overview.Sums, overview.SumsPerDuration

#### func (*Client) Average

```go
func (c *Client) Average(name string) (float64, int64, error)
```
Average returns the average value (float64) and number of cases (int64) of a
named server-side Average metric; or a non-nil error.

#### func (*Client) AveragePerDuration

```go
func (c *Client) AveragePerDuration(name string) (float64, int64, time.Time, error)
```
AveragePerDuration returns the average (float64), number of cases (int64) and
the until-timestamp (time.Time) of a named server-side AveragePerDuration
metric; or a non-nil error.

#### func (*Client) Close

```go
func (c *Client) Close()
```
Close disconnects the client. It may be deferred, similar to file closing.

#### func (*Client) Count

```go
func (c *Client) Count(name string) (int64, error)
```
Count returns the number of observations (int64) of a named server-side Count
metric; or a non-nil error.

#### func (*Client) CountPerDuration

```go
func (c *Client) CountPerDuration(name string) (int64, time.Time, error)
```
CountPerDuration returns the number of observations (int64) and the
until-timestamp (time.Time) of a named server-side CountPerDuration metric; or a
non-nil error.

#### func (*Client) Sum

```go
func (c *Client) Sum(name string) (float64, int64, error)
```
Sum returns the sum of observations (float64) and number of cases (int32) of a
named server-side Sum metric; or a non-nil error.

#### func (*Client) SumPerDuration

```go
func (c *Client) SumPerDuration(name string) (float64, int64, time.Time, error)
```
SumPerDuration returns the sum of observations (float64), the number of cases
(int32) and the until-timestamp (time.Time) of a named server-side
SumPerDuration metric; or a non-nil error.
