package reporter

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/api"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	conn            *grpc.ClientConn
	client          api.ReporterClient
	backoffTries    int
	backoffInterval time.Duration
}

// NewClient returns an initialized client, or a non-nil error.
// The addr argument is the "ip:port" to connect to; "ip" being optional (defaults to
// localhost). The default backup policy (to overcome network errors) is retry up to 5
// times, with delays 50, 100, 150, 200, 250 and 300 milliseconds between tries.
// This can be changed using WithBackoffPolicy.
func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, rtmerror.NewError("failed to dial %q", addr).WithError(err)
	}

	client := api.NewReporterClient(conn)

	return &Client{
		conn:            conn,
		client:          client,
		backoffTries:    5,
		backoffInterval: 50 * time.Millisecond,
	}, nil
}

// WithBackoffPolicy overrules the default retry/backoff policy for network operations.
// The first argument specifies how many times should be retried, the second specifies
// the base wait time. The actual wait time is the retry number times his wait time; e.g.,
// when the base wait time is 30ms, then for 5 retries the actual wait time is
// 30, 60, 90, 120 and 150 milliseconds.
//
// Example:
//
//  // This retry policy is stubborn but goes easy on resources. It retries 10 times,
//  // but waits 1, 2, 3, etc. seconds between consecutive tries.
//  // So, the client will only return an error for Sum(), Average() etc. after 55
//  // seconds (1+2+3+...+10).
//  c, err := NewClient(":1234").WithBackoffPolicy(10, time.Second)
func (c *Client) WithBackoffPolicy(tries int, interval time.Duration) *Client {
	c.backoffTries = tries
	c.backoffInterval = interval
	return c
}

// Close disconnects the client. It may be deferred, similar to file closing.
func (c *Client) Close() {
	c.conn.Close()
}

// AllNamesReturn is returned by AllNames and lists the names of averages, averages per duration,
// etc.
type AllNamesReturn struct {
	Averages            []string // names of averages
	AveragesPerDuration []string // names of averages per duration
	Counts              []string // names of counts
	CountsPerDuration   []string // names of counts per duration
	Sums                []string // names of sums
	SumsPerDuration     []string // names of sums per duration
}

// AllNames returns a list of server-registered named metrics, or a non-nil error.
//
// Example:
//
//  overview, err := client.AllNames()
//  if err != nil { ... }
//  for _, n := range overview.Averages {
//    fmt.Printf("There is a named average metric called %s\n", n)
//  }
//  // and so on with overview.AveragesPerDuration, overview.Counts,
//  // overview.CountsPerDuration, overview.Sums, overview.SumsPerDuration
//
// See demo/client_allnames.go for a full example.
func (c *Client) AllNames() (*AllNamesReturn, error) {
	var err error
	var resp *api.AllNamesResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.AllNames(context.Background(), &api.EmptyRequest{})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return nil, retryableError("AllNames", err)
	}
	return &AllNamesReturn{
		Averages:            resp.GetAverageNames(),
		AveragesPerDuration: resp.GetAveragePerDurationNames(),
		Counts:              resp.GetCountNames(),
		CountsPerDuration:   resp.GetCountPerDurationNames(),
		Sums:                resp.GetSumNames(),
		SumsPerDuration:     resp.GetSumPerDurationNames(),
	}, nil
}

type averageDump struct {
	Name  string
	Value float64
	N     int64
}
type averagePerDurationDump struct {
	Name  string
	Value float64
	N     int64
	Until time.Time
}
type countDump struct {
	Name  string
	Value int64
}
type countPerDurationDump struct {
	Name  string
	Value int64
	Until time.Time
}
type sumDump struct {
	Name  string
	Value float64
	N     int64
}
type sumPerDurationDump struct {
	Name  string
	Value float64
	N     int64
	Until time.Time
}

// FullDumpReturn is returned by FullDump and lists all averages, averages per duration, etc.,
// each including their name and reported state.
type FullDumpReturn struct {
	Averages            []averageDump
	AveragesPerDuration []averagePerDurationDump
	Counts              []countDump
	CountsPerDuration   []countPerDurationDump
	Sums                []sumDump
	SumsPerDuration     []sumPerDurationDump
}

// FullDump returns all names and states of server-known metrics. The return value is a
// (reference to a) structure with the fields:
//
//  Averages:            an array of structs with a Name (string), Value (float64), N (int64)
//  AveragesPerDuration: similar to the above, but the fields also have Until (time.Time)
//  Counts:              an array of structs with a Name (string), Value (int64)
//  CountsPerDuration:   similar to the above, but the fields also have Until (time.Time)
//  Sums:                an array of structs with a Name (string), Value (float64), N (int64)
//  SumsPerDuration:     similar to the above, but the fields also have Until (time.Time)
//
// See the package overview or demo/client_fulldump.go for a complete example.
func (c *Client) FullDump() (*FullDumpReturn, error) {
	var err error
	var resp *api.FullDumpResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.FullDump(context.Background(), &api.EmptyRequest{})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return nil, retryableError("FullDump", err)
	}

	ret := &FullDumpReturn{}

	for _, named := range resp.GetNamedAverages() {
		ret.Averages = append(ret.Averages, averageDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
			N:     named.GetN(),
		})
	}
	for _, named := range resp.GetNamedAveragesPerDuration() {
		ts, err := timeOf(named.GetUntil())
		if err != nil {
			return nil, err
		}
		ret.AveragesPerDuration = append(ret.AveragesPerDuration, averagePerDurationDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
			N:     named.GetN(),
			Until: ts,
		})
	}

	for _, named := range resp.GetNamedCounts() {
		ret.Counts = append(ret.Counts, countDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
		})
	}
	for _, named := range resp.GetNamedCountsPerDuration() {
		ts, err := timeOf(named.GetUntil())
		if err != nil {
			return nil, err
		}
		ret.CountsPerDuration = append(ret.CountsPerDuration, countPerDurationDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
			Until: ts,
		})
	}

	for _, named := range resp.GetNamedSums() {
		ret.Sums = append(ret.Sums, sumDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
			N:     named.GetN(),
		})
	}
	for _, named := range resp.GetNamedSumsPerDuration() {
		ts, err := timeOf(named.GetUntil())
		if err != nil {
			return nil, err
		}
		ret.SumsPerDuration = append(ret.SumsPerDuration, sumPerDurationDump{
			Name:  named.GetName(),
			Value: named.GetValue(),
			N:     named.GetN(),
			Until: ts,
		})
	}

	return ret, nil
}

// Average returns the average value (float64) and number of cases (int64) of a named
// server-side Average metric; or a non-nil error.
func (c *Client) Average(name string) (float64, int64, error) {
	var err error
	var resp *api.AverageResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.Average(context.Background(), &api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return 0.0, int64(0), retryableError("Average", err)
	}
	return resp.GetAverage(), resp.GetN(), nil
}

// AveragePerDuration returns the average (float64), number of cases (int64) and the until-timestamp
// (time.Time) of a named server-side AveragePerDuration metric; or a non-nil error.
func (c *Client) AveragePerDuration(name string) (float64, int64, time.Time, error) {
	var err error
	var resp *api.AveragePerDurationResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.AveragePerDuration(context.Background(),
			&api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return 0.0, int64(0), time.Now(), retryableError("AveragePerDuration", err)
	}
	ts, err := timeOf(resp.GetUntil())
	if err != nil {
		return 0.0, int64(0), time.Now(), err
	}
	return resp.GetAverage(), resp.GetN(), ts, nil
}

// Count returns the number of observations (int64) of a named
// server-side Count metric; or a non-nil error.
func (c *Client) Count(name string) (int64, error) {
	var err error
	var resp *api.CountResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.Count(context.Background(), &api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return int64(0), retryableError("Count", err)
	}
	return resp.GetCount(), nil
}

// CountPerDuration returns the number of observations (int64) and the until-timestamp
// (time.Time) of a named server-side CountPerDuration metric; or a non-nil error.
func (c *Client) CountPerDuration(name string) (int64, time.Time, error) {
	var err error
	var resp *api.CountPerDurationResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.CountPerDuration(context.Background(),
			&api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return int64(0), time.Now(), retryableError("CountPerDuration", err)
	}
	ts, err := timeOf(resp.GetUntil())
	if err != nil {
		return int64(0), time.Now(), err
	}
	return resp.GetCount(), ts, nil
}

// Sum returns the sum of observations (float64) and number of cases (int32) of a named
// server-side Sum metric; or a non-nil error.
func (c *Client) Sum(name string) (float64, int64, error) {
	var err error
	var resp *api.SumResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.Sum(context.Background(), &api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return 0.0, int64(0), retryableError("Sum", err)
	}
	return resp.GetSum(), resp.GetN(), nil
}

// SumPerDuration returns the sum of observations (float64), the number of cases (int32)
// and the until-timestamp (time.Time) of a named server-side SumPerDuration metric; or
// a non-nil error.
func (c *Client) SumPerDuration(name string) (float64, int64, time.Time, error) {
	var err error
	var resp *api.SumPerDurationResponse

	for i := 0; i < c.backoffTries; i++ {
		resp, err = c.client.SumPerDuration(context.Background(), &api.NameRequest{Name: name})
		if err == nil || i == c.backoffTries-1 {
			break
		}
		time.Sleep(c.backoffInterval * time.Duration(i))
	}
	if err != nil {
		return 0.0, int64(0), time.Now(), retryableError("SumPerDuration", err)
	}
	ts, err := timeOf(resp.GetUntil())
	if err != nil {
		return 0.0, int64(0), time.Now(), err
	}
	return resp.GetSum(), resp.GetN(), ts, nil
}

func retryableError(serviceName string, err error) error {
	return rtmerror.NewError("failed in service %q", serviceName).
		WithError(err).WithRetryable(true)
}

func timeOf(t *timestamp.Timestamp) (time.Time, error) {
	ts, err := ptypes.Timestamp(t)
	if err != nil {
		return ts, rtmerror.NewError("timestamp conversion failed").
			WithError(err)
	}
	return ts, nil
}
