package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateLineRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 1200},
			Height: &wrapperspb.Int32Value{Value: 700},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 20},
			MarginBottom: &wrapperspb.Int32Value{Value: 70},
			MarginLeft:   &wrapperspb.Int32Value{Value: 40},
			MarginRight:  &wrapperspb.Int32Value{Value: 30},
		},
		Axes: &render.ChartAxes{
			AxisBottom: &render.ChartScale{
				Kind: render.ChartScale_BAND,
				Domain: &render.ChartScale_DomainCategories{
					DomainCategories: &render.DomainCategories{
						Categories: []string{"a1", "a2", "a3", "a4", "a5", "a6"},
					},
				},
				NoBoundariesOffset: true,
				InnerPadding:       &wrapperspb.FloatValue{Value: 0.0},
				OuterPadding:       &wrapperspb.FloatValue{Value: 0.0},
			},
			AxisLeft: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   200,
					},
				},
			},
		},
		Views: []*render.ChartView{
			{
				Kind: render.ChartView_LINE,
				Values: &render.ChartView_ScalarValues{
					ScalarValues: &render.ChartViewScalarValues{
						Values: []float32{12, 100, 120, 180, 40, 8},
					},
				},
				PointLabelPosition: render.ChartView_TOP_RIGHT,
			},
		},
	}
}
