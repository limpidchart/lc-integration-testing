package integration_test

import (
	"context"
	"encoding/base64"
	"testing"
	"time"
)

func TestCreateChart_HTTP(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name          string
		requestPath   string
		replyDataPath string
	}{
		{
			"area",
			createAreaRequestJSONPath,
			createAreaReplyData,
		},
		{
			"horizontal_bar",
			createHorizontalBarRequestJSONPath,
			createHorizontalBarReplyData,
		},
		{
			"line_and_vertical_bar",
			createLineAndVerticalBarRequestJSONPath,
			createLineAndVerticalBarReplyData,
		},
		{
			"line",
			createLineRequestJSONPath,
			createLineReplyData,
		},
		{
			"scatter",
			createScatterRequestJSONPath,
			createScatterReplyData,
		},
		{
			"stacked_horizontal_bar",
			createStackedHorizontalBarRequestJSONPath,
			createStackedHorizontalBarReplyData,
		},
		{
			"stacked_vertical_bar",
			createStackedVerticalBarRequestJSONPath,
			createStackedVerticalBarReplyData,
		},
		{
			"two_lines",
			createTwoLinesRequestJSONPath,
			createTwoLinesReplyData,
		},
		{
			"two_scatters",
			createTwoScattersRequestJSONPath,
			createTwoScattersReplyData,
		},
		{
			"vertical_bar",
			createVerticalBarRequestJSONPath,
			createVerticalBarReplyData,
		},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testStart := time.Now().UTC()

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(LCAPIHTTPTimeout)*time.Second)
			defer cancel()

			chartRep, err := createChartAndParseGoodReplyHTTP(ctx, tc.requestPath)
			if err != nil {
				t.Fatalf("Unable to create chart: %s", err)
			}

			if errBasicFields := checkBasicCreateChartReplyFields(chartRep, testStart); errBasicFields != nil {
				t.Fatalf("Unable to validate reply basic fields: %s", errBasicFields)
			}

			if errTSFields := checkCreateChartReplyCreatedAtAndDeletedAtEqual(chartRep); errTSFields != nil {
				t.Fatalf("Unable to validate reply timestamp fields: %s", errTSFields)
			}

			actualData, err := base64.StdEncoding.DecodeString(chartRep.ChartData)
			if err != nil {
				t.Fatalf("Unable to decode base64 chart data: %s", err)
			}

			if err := compareExpectedAndActualChartLines(tc.replyDataPath, actualData); err != nil {
				t.Fatalf("Unable to validate reply chart data: %s", err)
			}
		})
	}
}
