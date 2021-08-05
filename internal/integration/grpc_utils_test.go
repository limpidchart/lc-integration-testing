package integration_test

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const LCAPIgRPCTimeout = 60

const (
	lcAPIgRPCAddress = "localhost:54010"
)

func createChartAndParseGoodReplyGRPC(ctx context.Context, req *render.CreateChartRequest) (*render.ChartReply, error) {
	chartAPIServerConn, err := grpc.DialContext(ctx, lcAPIgRPCAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to lc-api server: %w", err)
	}

	chartAPIClient := render.NewChartAPIClient(chartAPIServerConn)

	reply, err := chartAPIClient.CreateChart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to create chart via gRPC: %w", err)
	}

	return reply, nil
}
