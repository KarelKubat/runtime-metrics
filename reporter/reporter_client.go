package reporter

import (
	"fmt"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client ReporterClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to dial %q: %v", addr, err)
	}

	client := NewReporterClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

type AllNamesReturn struct {
	Averages            []string
	AveragesPerDuration []string
	Counters            []string
	CountersPerDuration []string
	Sums                []string
	SumsPerDuration     []string
}

func (c *Client) AllNames() (*AllNamesReturn, error) {
	resp, err := c.client.AllNames(context.Background(), &AllNamesRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed at AllNames service: %v", err)
	}
	return &AllNamesReturn{
		Averages:            resp.GetAverageNames(),
		AveragesPerDuration: resp.GetAveragePerDurationNames(),
		Counters:            resp.GetCounterNames(),
		CountersPerDuration: resp.GetCounterPerDurationNames(),
		Sums:                resp.GetSumNames(),
		SumsPerDuration:     resp.GetSumPerDurationNames(),
	}, nil
}

func (c *Client) Average(name string) (float64, int64, error) {
	resp, err := c.client.Average(&NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), fmt.Errorf("failed at Average service: %v", err)
	}
	return resp.GetAverage(), resp.GetN(), nil
}

func (c *Client) AveragePerDuration(name string) (float64, int64, time.Time, error) {
	resp, err := c.client.AveragePerDuration(&NameRequest{Name: name})
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("failed at Average service: %v", err)
	}
	ts, err := timestamp.Timestamp(resp.GetUntil())
	if err != nil {
		return 0.0, int64(0), time.Now(), fmt.Errorf("failed to convert timestamp: %v", err)
	}
	return resp.GetAverage(), resp.GetN(), ts, nil
}
