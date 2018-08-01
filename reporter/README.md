# reporter
--
    import "github.com/KarelKubat/runtime-metrics/reporter"

reporter implements a metrics reporting server and client.


### The Server

The server is started using reporter.StartReporter(addr), where addr defines the
IP and port to listen, separated by :. The IP can be left out. Example:

    err := reporter.StartReporter(":9000")    // bind to TCP port 9000
    if err != nil { ... }                     // probably port already taken

Typically this will be wrapped in a go function, so that the reporting server
runs in its own thread:

    go func() {
      if err := reporter.StartReporter(":9000"); err != nil {
        ... // probably port 9000 is already taken
      }
    }()

Any metrics that are registered are available for the server to be published:

    func checkErr(err) {
      if err != nil { ... } // name collision
    }
    ...
    // Create and register metrics
    myAverage := base.NewAverage()
    checkErr(registry.AddAverage("average-metric", myAverage))
    mySumPD := base.NewSumPerDuration(time.Minute)
    checkErr(registry.AddSumPerDuration("sum-per-minute", mySumPD))
    ...
    // As metrics are updated, the reporting server will publish them.
    myAverage.Mark(3.14)
    mySumPD.Mark(2.71)


### The Client

The client is instantiated using reporter.NewClient(addr), where again addr is a
string defining the IP address and TCP port:

    c, err := reporter.NewClient(":9000")
    checkErr(err)
    defer c.Close()

The client has a number of handy methods that contact the client to discover the
names of metrics or to fetch metric values:

    allNames, err := c.AllNames()
    // allNames.Averages           is an array of strings (names) of all Average-type metrics
    // allNames.AveragesByDuration is an array of strings (names) of all AveragePerDuration-type metrics
    // allNames.Counts             is an array of strings (names) of all Count-type metrics
    // allNames.CountsPerDuration  is an array of strings (names) of all CountPerDuration-type metrics
    // allNames.Sums               is an array of strings (names) of all Sum-type metrics
    // allNames.SumsPerDuration    is an array of strings (names) of all SumPerDuration-type metrics

In order to fetch the values of a metric, the client calls c.Average(name),
c.Sum(name) etc. The returned values are always what the base type returns, and
an error:

    avg, n, err := c.Average("my-average")
    // val is the average
    // n   is the number of observations
    // err is nil or an error

A full dump of all server-known metrics can be obtained using FullDump(). See
also demo/demosrc/client_fulldump.go for an example.

    dump, err := c.FullDump()
    if err != nil { ... }
    for _, av := range dump.Averages {
      // av.Name is the name, av.Value is the average, av.N is the number of cases
    }
    for _, avPD := range dump.AveragesPerDuration {
      // avPD.Name is the name, avPD.Value is the average, avPD.N is the number of cases,
      // av.Until is the up-to timestamp
    }
    for _, c := range dump.Counts {
      // c.Name is the name, c.Value is the count
    }
    for _, cPD := range dump.CountsPerDuration {
      // cPD.Name is the name, cPD.Value is the count, cPD.Until is the up-to timestamp
    }
    for _, s := range dump.Sums {
      // s.Name is the name, s.Value is the sum, s.N is the number of cases
    }
    for _, sPD := range dump.SumsPerDuration {
      // sPD.Name is the name, sPD.Value is the sum, sPD.N is the number of cases,
      // sPD.Until is the up-to timestamp
    }

The network calls that the client issues to obtain metrics, are subject to a
retry policy. The default policy is that if the network call fails, then the
client waits for 50 milliseconds and retries. If that fails, the server waits
for 100 milliseconds and retries again. If that fails, then the wait time is
extended by another 50 milliseconds (becoming 150), and the client retries
again.

This policy is defined by two numbers: the allowed retries (defaults to 5) and
the duration by which the wait time is extended each time that a call fails
(defaults to 50 milliseconds). This backoff policy can be overruled when
constructing a client using WithBackoffPolicy(), for example:

    c, err := reporter.NewClient(":9000").WithBackoffPolicy(
      10,                           // retry up to 10 times
      100 * time.Millisecond)       // 100ms between the first failed call and the first retry,
                                    // 1s between the 9th and 10th retry

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
is the "ip:port" to connect to; "ip" being optional (defaults to localhost). The
default backup policy (to overcome network errors) is retry up to 5 times, with
delays 50, 100, 150, 200, 250 and 300 milliseconds between tries. This can be
changed using WithBackoffPolicy.

#### func (*Client) AllNames

```go
func (c *Client) AllNames() (*AllNamesReturn, error)
```
AllNames returns a list of server-registered named metrics, or a non-nil error.

Example:

    overview, err := client.AllNames()
    if err != nil { ... }
    for _, n := range overview.Averages {
      fmt.Printf("There is a named average metric called %s\n", n)
    }
    // and so on with overview.AveragesPerDuration, overview.Counts,
    // overview.CountsPerDuration, overview.Sums, overview.SumsPerDuration

See demo/demosrc/client_allnames.go for a full example.

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

#### func (*Client) FullDump

```go
func (c *Client) FullDump() (*FullDumpReturn, error)
```
FullDump returns all names and states of server-known metrics. The return value
is a (reference to a) structure with the fields:

    Averages:            an array of structs with a Name (string), Value (float64), N (int64)
    AveragesPerDuration: similar to the above, but the fields also have Until (time.Time)
    Counts:              an array of structs with a Name (string), Value (int64)
    CountsPerDuration:   similar to the above, but the fields also have Until (time.Time)
    Sums:                an array of structs with a Name (string), Value (float64), N (int64)
    SumsPerDuration:     similar to the above, but the fields also have Until (time.Time)

See the package overview or demo/demosrc/client_fulldump.go for a complete
example.

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

#### func (*Client) WithBackoffPolicy

```go
func (c *Client) WithBackoffPolicy(tries int, interval time.Duration) *Client
```
WithBackoffPolicy overrules the default retry/backoff policy for network
operations. The first argument specifies how many times should be retried, the
second specifies the base wait time. The actual wait time is the retry number
times his wait time; e.g., when the base wait time is 30ms, then for 5 retries
the actual wait time is 30, 60, 90, 120 and 150 milliseconds.

Example:

    // This retry policy is stubborn but goes easy on resources. It retries 10 times,
    // but waits 1, 2, 3, etc. seconds between consecutive tries.
    // So, the client will only return an error for Sum(), Average() etc. after 55
    // seconds (1+2+3+...+10).
    c, err := NewClient(":1234").WithBackoffPolicy(10, time.Second)

#### type FullDumpReturn

```go
type FullDumpReturn struct {
	Averages            []averageDump
	AveragesPerDuration []averagePerDurationDump
	Counts              []countDump
	CountsPerDuration   []countPerDurationDump
	Sums                []sumDump
	SumsPerDuration     []sumPerDurationDump
}
```

FullDumpReturn is returned by FullDump and lists all averages, averages per
duration, etc., each including their name and reported state.
