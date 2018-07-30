package reporter

import (
	"fmt"
	"net"

	"github.com/KarelKubat/runtime-metrics/api"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
}

// AllNames uses the API to return registered named metrics.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) AllNames(ctx context.Context, in *api.AllNamesRequest) (*api.AllNamesResponse, error) {
	return &api.AllNamesResponse{
		AverageNames:            registry.AverageNames(),
		AveragePerDurationNames: registry.AveragePerDurationNames(),
		CountNames:              registry.CountNames(),
		CountPerDurationNames:   registry.CountPerDurationNames(),
		SumNames:                registry.SumNames(),
		SumPerDurationNames:     registry.SumPerDurationNames(),
	}, nil
}

// Average uses the API to return the state of a named Average.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) Average(ctx context.Context, in *api.NameRequest) (*api.AverageResponse, error) {
	av, err := registry.GetAverage(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n := av.Report()
	return &api.AverageResponse{
		Average: val,
		N:       n,
	}, nil
}

// Average uses the API to return the state of a named AveragePerDuration.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) AveragePerDuration(ctx context.Context, in *api.NameRequest) (*api.AveragePerDurationResponse, error) {
	av, err := registry.GetAveragePerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &api.AveragePerDurationResponse{
		Average: val,
		N:       n,
		Until:   ts,
	}, nil
}

// Average uses the API to return the state of a named Count.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) Count(ctx context.Context, in *api.NameRequest) (*api.CountResponse, error) {
	av, err := registry.GetCount(in.GetName())
	if err != nil {
		return nil, err
	}
	return &api.CountResponse{
		Count: av.Report(),
	}, nil
}

// Average uses the API to return the state of a named CountPerDuration.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) CountPerDuration(ctx context.Context, in *api.NameRequest) (*api.CountPerDurationResponse, error) {
	av, err := registry.GetCountPerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &api.CountPerDurationResponse{
		Count: val,
		Until: ts,
	}, nil
}

// Average uses the API to return the state of a named Sum.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) Sum(ctx context.Context, in *api.NameRequest) (*api.SumResponse, error) {
	av, err := registry.GetSum(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n := av.Report()
	return &api.SumResponse{
		Sum: val,
		N:   n,
	}, nil
}

// Average uses the API to return the state of a named SumPerDuration.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) SumPerDuration(ctx context.Context, in *api.NameRequest) (*api.SumPerDurationResponse, error) {
	av, err := registry.GetSumPerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &api.SumPerDurationResponse{
		Sum:   val,
		N:     n,
		Until: ts,
	}, nil
}

// StartReporter starts the reporting server or returns a non-nil error.
// The argument is an address in the format "ip:port", where "ip" is
// optional. This is the binding address.
func StartReporter(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen to address %q: %v", addr, err)
	}

	s := &server{}
	grpcServer := grpc.NewServer()
	api.RegisterReporterServer(grpcServer, s)
	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	// notreached
	return nil
}
