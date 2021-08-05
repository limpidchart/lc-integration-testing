package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateTwoScattersRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 800},
			Height: &wrapperspb.Int32Value{Value: 600},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 90},
			MarginBottom: &wrapperspb.Int32Value{Value: 50},
			MarginLeft:   &wrapperspb.Int32Value{Value: 60},
			MarginRight:  &wrapperspb.Int32Value{Value: 40},
		},
		Axes: &render.ChartAxes{
			AxisLeft: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   100,
					},
				},
			},
			AxisRight: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   100,
					},
				},
			},
			AxisTop: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   200.0,
					},
				},
			},
			AxisBottom: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   200.0,
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
								X: 20,
								Y: 90,
							},
							{
								X: 12,
								Y: 54,
							},
							{
								X: 25,
								Y: 70,
							},
							{
								X: 33,
								Y: 40,
							},
						},
					},
				},
				Colors: &render.ChartViewColors{
					PointFill: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#808080",
						},
					},
					PointStroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#000000",
						},
					},
				},
			},
			{
				Kind: render.ChartView_SCATTER,
				Values: &render.ChartView_PointsValues{
					PointsValues: &render.ChartViewPointsValues{
						Points: []*render.ChartViewPointsValues_Point{
							{
								X: 120,
								Y: 10,
							},
							{
								X: 143,
								Y: 34,
							},
							{
								X: 170,
								Y: 14,
							},
							{
								X: 190,
								Y: 13,
							},
						},
					},
				},
				Colors: &render.ChartViewColors{
					PointFill: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#000000",
						},
					},
					PointStroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#808080",
						},
					},
				},
			},
		},
	}
}
