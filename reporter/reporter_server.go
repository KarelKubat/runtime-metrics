package reporter

import (
	"fmt"
	"net"

	"github.com/KarelKubat/runtime-metrics/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) AllNames(ctx context.Context, in *AllNamesRequest) (*AllNamesResponse, error) {
	return &DiscoverResponse{
		AverageNames:            registry.AverageNames(),
		AveragePerDurationNames: registry.AveragePerDurationNames(),
		CounterNames:            registry.CounterNames(),
		CounterPerDurationNames: registry.CounterPerDurationNames(),
		SumNames:                registry.SumNames(),
		SumPerDurationNames:     registry.SumPerDurationNames(),
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
