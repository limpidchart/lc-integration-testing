package fixtures

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateLineAndVerticalBarRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Title: "Cost of living index in Berlin",
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
			AxisBottom: &render.ChartScale{
				Kind:       render.ChartScale_BAND,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 700},
				Domain: &render.ChartScale_DomainCategories{
					DomainCategories: &render.DomainCategories{
						Categories: []string{
							"2015",
							"2016",
							"2017",
							"2018",
							"2019",
							"2020",
							"2021",
						},
					},
				},
			},
			AxisLeft: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 460},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   85,
					},
				},
			},
		},
		Views: []*render.ChartView{
			{
				Kind: render.ChartView_VERTICAL_BAR,
				Values: &render.ChartView_BarsValues{
					BarsValues: &render.ChartViewBarsValues{
						BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
							{
								Values: []float32{74.72, 66.34, 64.18, 74.32, 68.62, 64.76, 72.52},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#77ab59",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#36802d",
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Kind: render.ChartView_LINE,
				Values: &render.ChartView_ScalarValues{
					ScalarValues: &render.ChartViewScalarValues{
						Values: []float32{74.72, 66.34, 64.18, 74.32, 68.62, 64.76, 72.52},
					},
				},
				PointLabelVisible: &wrapperspb.BoolValue{Value: false},
				PointType:         render.ChartView_CIRCLE,
				Colors: &render.ChartViewColors{
					Stroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#234d20",
						},
					},
					PointStroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#234d20",
						},
					},
					PointFill: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#234d20",
						},
					},
				},
			},
		},
	}
}
