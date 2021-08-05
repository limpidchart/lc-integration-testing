package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/limpidchart/lc-integration-testing/internal/integration/fixtures"
	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func TestCreateChart_gRPC(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name          string
		request       *render.CreateChartRequest
		replyDataPath string
	}{
		{
			"area",
			fixtures.CreateAreaRequest(),
			createAreaReplyData,
		},
		{
			"horizontal_bar",
			fixtures.CreateHorizontalBarRequest(),
			createHorizontalBarReplyData,
		},
		{
			"line_and_vertical_bar",
			fixtures.CreateLineAndVerticalBarRequest(),
			createLineAndVerticalBarReplyData,
		},
		{
			"line",
			fixtures.CreateLineRequest(),
			createLineReplyData,
		},
		{
			"scatter",
			fixtures.CreateScatterRequest(),
			createScatterReplyData,
		},
		{
			"stacked_horizontal_bar",
			fixtures.CreateStackedHorizontalBarRequest(),
			createStackedHorizontalBarReplyData,
		},
		{
			"stacked_vertical_bar",
			fixtures.CreateStackedVerticalBarRequest(),
			createStackedVerticalBarReplyData,
		},
		{
			"two_lines",
			fixtures.CreateTwoLinesRequest(),
			createTwoLinesReplyData,
		},
		{
			"two_scatters",
			fixtures.CreateTwoScattersRequest(),
			createTwoScattersReplyData,
		},
		{
			"vertical_bar",
			fixtures.CreateVerticalBarRequest(),
			createVerticalBarReplyData,
		},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testStart := time.Now().UTC()

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(LCAPIgRPCTimeout)*time.Second)
			defer cancel()

			chartRepProto, err := createChartAndParseGoodReplyGRPC(ctx, tc.request)
			if err != nil {
				t.Fatalf("Unable to create chart: %s", err)
			}

			chartRep := chartReplyFromProtobuf(chartRepProto)

			if errBasicFields := checkBasicCreateChartReplyFields(chartRep, testStart); errBasicFields != nil {
				t.Fatalf("Unable to validate reply basic fields: %s", errBasicFields)
			}

			if errTSFields := checkCreateChartReplyCreatedAtAndDeletedAtEqual(chartRep); errTSFields != nil {
				t.Fatalf("Unable to validate reply timestamp fields: %s", errTSFields)
			}

			if err := compareExpectedAndActualChartLines(tc.replyDataPath, []byte(chartRep.ChartData)); err != nil {
				t.Fatalf("Unable to validate reply chart data: %s", err)
			}
		})
	}
}
