package reporter

import (
	"fmt"
	"time"

	"github.com/KarelKubat/runtime-metrics/api"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client api.ReporterClient
}

// NewClient returns an initialized client, or a non-nil error.
// The addr argument is the "ip:port" to connect to; "ip" being optional (defaults to
// localhost).
func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to dial %q: %v", addr, err)
	}

	client := api.NewReporterClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
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
func (c *Client) AllNames() (*AllNamesReturn, error) {
	resp, err := c.client.AllNames(context.Background(), &api.EmptyRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed at AllNames service: %v", err)
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
// Averages: an array of structs with a Name (string), Value (float64), N (int64)
// AveragesPerDuration: similar to the above, but the fields also have Until (time.Time)
// Counts: an array of structs with a Name (string), Value (int64)
// CountsPerDuration: similar to the above, but the fields also have Until (time.Time)
// Sums: an array of structs with a Name (string), Value (float64), N (int64)
// SumsPerDuration: similar to the above, but the fields also have Until (time.Time)
func (c *Client) FullDump() (*FullDumpReturn, error) {
	resp, err := c.client.FullDump(context.Background(), &api.EmptyRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed at FullDump service: %v", err)
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
		ts, err := ptypes.Timestamp(named.GetUntil())
		if err != nil {
			return nil, fmt.Errorf("timestamp conversion failed: %v", err)
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
		ts, err := ptypes.Timestamp(named.GetUntil())
		if err != nil {
			return nil, fmt.Errorf("timestamp conversion failed: %v", err)
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
		ts, err := ptypes.Timestamp(named.GetUntil())
		if err != nil {
			return nil, fmt.Errorf("timestamp conversion failed: %v", err)
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
	resp, err := c.client.Average(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), fmt.Errorf("failed at Average service: %v", err)
	}
	return resp.GetAverage(), resp.GetN(), nil
}

// AveragePerDuration returns the average (float64), number of cases (int64) and the until-timestamp
// (time.Time) of a named server-side AveragePerDuration metric; or a non-nil error.
func (c *Client) AveragePerDuration(name string) (float64, int64, time.Time, error) {
	resp, err := c.client.AveragePerDuration(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("failed at AveragePerDuration service: %v", err)
	}
	ts, err := ptypes.Timestamp(resp.GetUntil())
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return resp.GetAverage(), resp.GetN(), ts, nil
}

// Count returns the number of observations (int64) of a named
// server-side Count metric; or a non-nil error.
func (c *Client) Count(name string) (int64, error) {
	resp, err := c.client.Count(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return int64(0), fmt.Errorf("failed at Count service: %v", err)
	}
	return resp.GetCount(), nil
}

// CountPerDuration returns the number of observations (int64) and the until-timestamp
// (time.Time) of a named server-side CountPerDuration metric; or a non-nil error.
func (c *Client) CountPerDuration(name string) (int64, time.Time, error) {
	resp, err := c.client.CountPerDuration(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return int64(0), time.Now(), fmt.Errorf("failed at CountPerDuration service: %v", err)
	}
	ts, err := ptypes.Timestamp(resp.GetUntil())
	if err != nil {
		return int64(0), time.Now(), fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return resp.GetCount(), ts, nil
}

// Sum returns the sum of observations (float64) and number of cases (int32) of a named
// server-side Sum metric; or a non-nil error.
func (c *Client) Sum(name string) (float64, int64, error) {
	resp, err := c.client.Sum(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), fmt.Errorf("failed at Sum service: %v", err)
	}
	return resp.GetSum(), resp.GetN(), nil
}

// SumPerDuration returns the sum of observations (float64), the number of cases (int32)
// and the until-timestamp (time.Time) of a named server-side SumPerDuration metric; or
// a non-nil error.
func (c *Client) SumPerDuration(name string) (float64, int64, time.Time, error) {
	resp, err := c.client.SumPerDuration(context.Background(), &api.NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("failed at SumPerDuration service: %v", err)
	}
	ts, err := ptypes.Timestamp(resp.GetUntil())
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return resp.GetSum(), resp.GetN(), ts, nil
}
