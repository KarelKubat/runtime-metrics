package reporter

import (
	"fmt"
	"net"

	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) AllNames(ctx context.Context, in *AllNamesRequest) (*AllNamesResponse, error) {
	return &AllNamesResponse{
		AverageNames:            registry.AverageNames(),
		AveragePerDurationNames: registry.AveragePerDurationNames(),
		CounterNames:            registry.CounterNames(),
		CounterPerDurationNames: registry.CounterPerDurationNames(),
		SumNames:                registry.SumNames(),
		SumPerDurationNames:     registry.SumPerDurationNames(),
	}, nil
}

func (s *server) Average(ctx context.Context, in *NameRequest) (*AverageResponse, error) {
	av, err := registry.GetAverage(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n := av.Report()
	return &AverageResponse{
		Average: val,
		N:       n,
	}, nil
}

func (s *server) AveragePerDuration(ctx context.Context, in *NameRequest) (*AveragePerDurationResponse, error) {
	av, err := registry.GetAveragePerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &AveragePerDurationResponse{
		Average: val,
		N:       n,
		Until:   ts,
	}, nil
}

func (s *server) Count(ctx context.Context, in *NameRequest) (*CountResponse, error) {
	av, err := registry.GetCounter(in.GetName())
	if err != nil {
		return nil, err
	}
	return &CountResponse{
		Count: av.Report(),
	}, nil
}

func (s *server) CountPerDuration(ctx context.Context, in *NameRequest) (*CountPerDurationResponse, error) {
	av, err := registry.GetCounterPerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &CountPerDurationResponse{
		Count: val,
		Until: ts,
	}, nil
}

func (s *server) Sum(ctx context.Context, in *NameRequest) (*SumResponse, error) {
	av, err := registry.GetSum(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n := av.Report()
	return &SumResponse{
		Sum: val,
		N:   n,
	}, nil
}

func (s *server) SumPerDuration(ctx context.Context, in *NameRequest) (*SumPerDurationResponse, error) {
	av, err := registry.GetSumPerDuration(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, err := ptypes.TimestampProto(until)
	if err != nil {
		return nil, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return &SumPerDurationResponse{
		Sum:   val,
		N:     n,
		Until: ts,
	}, nil
}

func StartReporter(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen to address %q: %v", addr, err)
	}

	s := &server{}
	grpcServer := grpc.NewServer()
	RegisterReporterServer(grpcServer, s)
	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	// notreached
	return nil
}
