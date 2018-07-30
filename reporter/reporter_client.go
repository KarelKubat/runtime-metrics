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
//  overview, err := client.AllNameS()
//  if err != nil { ... }
//  for _, n := range overview.Averages {
//    fmt.Printf("There is a named average metric called %s\n", n)
//  }
//  // and so on with overview.AveragesPerDuration, overview.Counts,
//  // overview.CountsPerDuration, overview.Sums, overview.SumsPerDuration
func (c *Client) AllNames() (*AllNamesReturn, error) {
	resp, err := c.client.AllNames(context.Background(), &api.AllNamesRequest{})
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
