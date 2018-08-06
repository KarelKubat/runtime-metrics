package reporter

import (
	"fmt"
	"net"
	"time"

	"github.com/KarelKubat/runtime-metrics/api"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/rtmerror"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
}

// AllNames uses the API to return registered named metrics.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) AllNames(ctx context.Context, in *api.EmptyRequest) (*api.AllNamesResponse, error) {
	return &api.AllNamesResponse{
		AverageNames:            registry.AverageNames(),
		AveragePerDurationNames: registry.AveragePerDurationNames(),
		CountNames:              registry.CountNames(),
		CountPerDurationNames:   registry.CountPerDurationNames(),
		SumNames:                registry.SumNames(),
		SumPerDurationNames:     registry.SumPerDurationNames(),
	}, nil
}

// FullDump uses the API to return all registered named metrics and their values.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) FullDump(ctx context.Context, in *api.EmptyRequest) (*api.FullDumpResponse, error) {
	ret := &api.FullDumpResponse{}

	for _, name := range registry.AverageNames() {
		metric, err := registry.AverageBy(name)
		if err != nil {
			return nil, err
		}
		val, n := metric.Report()
		ret.NamedAverages = append(ret.NamedAverages, &api.NamedAverage{
			Name:  name,
			Value: val,
			N:     n,
		})
	}

	for _, name := range registry.AveragePerDurationNames() {
		metric, err := registry.AveragePerDurationBy(name)
		if err != nil {
			return nil, err
		}
		val, n, until := metric.Report()
		ts, stderr := timestampOf(until)
		if stderr != nil {
			return nil, stderr
		}
		ret.NamedAveragesPerDuration = append(
			ret.NamedAveragesPerDuration, &api.NamedAveragePerDuration{
				Name:  name,
				Value: val,
				N:     n,
				Until: ts,
			})
	}

	for _, name := range registry.CountNames() {
		metric, err := registry.CountBy(name)
		if err != nil {
			return nil, err
		}
		val := metric.Report()
		ret.NamedCounts = append(ret.NamedCounts, &api.NamedCount{
			Name:  name,
			Value: val,
		})
	}

	for _, name := range registry.CountPerDurationNames() {
		metric, err := registry.CountPerDurationBy(name)
		if err != nil {
			return nil, err
		}
		val, until := metric.Report()
		ts, stderr := timestampOf(until)
		if stderr != nil {
			return nil, stderr
		}
		ret.NamedCountsPerDuration = append(
			ret.NamedCountsPerDuration, &api.NamedCountPerDuration{
				Name:  name,
				Value: val,
				Until: ts,
			})
	}

	for _, name := range registry.SumNames() {
		metric, err := registry.SumBy(name)
		if err != nil {
			return nil, err
		}
		val, n := metric.Report()
		ret.NamedSums = append(ret.NamedSums, &api.NamedSum{
			Name:  name,
			Value: val,
			N:     n,
		})
	}

	for _, name := range registry.SumPerDurationNames() {
		metric, err := registry.SumPerDurationBy(name)
		if err != nil {
			return nil, err
		}
		val, n, until := metric.Report()
		ts, stderr := timestampOf(until)
		if stderr != nil {
			return nil, stderr
		}
		ret.NamedSumsPerDuration = append(
			ret.NamedSumsPerDuration, &api.NamedSumPerDuration{
				Name:  name,
				Value: val,
				N:     n,
				Until: ts,
			})
	}

	return ret, nil
}

// Average uses the API to return the state of a named Average.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) Average(ctx context.Context, in *api.NameRequest) (*api.AverageResponse, error) {
	av, err := registry.AverageBy(in.GetName())
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
func (s *server) AveragePerDuration(ctx context.Context, in *api.NameRequest) (
	*api.AveragePerDurationResponse, error) {

	av, err := registry.AveragePerDurationBy(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, stderr := timestampOf(until)
	if stderr != nil {
		return nil, stderr
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
	av, err := registry.CountBy(in.GetName())
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
	av, err := registry.CountPerDurationBy(in.GetName())
	if err != nil {
		return nil, err
	}
	val, until := av.Report()
	ts, stderr := timestampOf(until)
	if stderr != nil {
		return nil, stderr
	}
	return &api.CountPerDurationResponse{
		Count: val,
		Until: ts,
	}, nil
}

// Average uses the API to return the state of a named Sum.
// This is wrapped in the client, this function isn't for public consumption.
func (s *server) Sum(ctx context.Context, in *api.NameRequest) (*api.SumResponse, error) {
	av, err := registry.SumBy(in.GetName())
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
	av, err := registry.SumPerDurationBy(in.GetName())
	if err != nil {
		return nil, err
	}
	val, n, until := av.Report()
	ts, stderr := timestampOf(until)
	if stderr != nil {
		return nil, stderr
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
		return rtmerror.NewError("failed to listen to address %q", addr).
			WithError(err)
	}

	s := &server{}
	grpcServer := grpc.NewServer()
	api.RegisterReporterServer(grpcServer, s)
	if err = grpcServer.Serve(lis); err != nil {
		return rtmerror.NewError("failed to start server").WithError(err)
	}

	// notreached
	return nil
}

func timestampOf(d time.Time) (*timestamp.Timestamp, error) {
	ts, err := ptypes.TimestampProto(d)
	if err != nil {
		return ts, fmt.Errorf("timestamp conversion failed: %v", err)
	}
	return ts, nil
}
