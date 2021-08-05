package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateScatterRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 800},
			Height: &wrapperspb.Int32Value{Value: 600},
		},
		Axes: &render.ChartAxes{
			AxisLeft: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 460},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   100.0,
					},
				},
			},
			AxisRight: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 460},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   100.0,
					},
				},
			},
			AxisTop: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 700},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   200,
					},
				},
			},
			AxisBottom: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 700},
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
				Kind: render.ChartView_SCATTER,
				Values: &render.ChartView_PointsValues{
					PointsValues: &render.ChartViewPointsValues{
						Points: []*render.ChartViewPointsValues_Point{
							{
								X: 20.1,
								Y: 54.11,
							},
							{
								X: 70.2,
								Y: 40.22,
							},
							{
								X: 130.3,
								Y: 50.33,
							},
							{
								X: 170.4,
								Y: 70.44,
							},
							{
								X: 20.5,
								Y: 90.55,
							},
							{
								X: 95.6,
								Y: 40.66,
							},
							{
								X: 130.7,
								Y: 12.77,
							},
							{
								X: 170.8,
								Y: 2.88,
							},
						},
					},
				},
				PointLabelPosition: render.ChartView_BOTTOM_RIGHT,
			},
		},
	}
}
