package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateStackedVerticalBarRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Title: "Cost of living index",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 800},
			Height: &wrapperspb.Int32Value{Value: 1000},
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
							"Russia",
							"Germany",
							"Netherlands",
							"Canada",
							"United States",
							"Australia",
						},
					},
				},
			},
			AxisLeft: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 860},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   450.0,
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
								Values: []float32{33.17, 70.62, 78.64, 70.08, 71.92, 84.14},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#00fff9",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#00a2c5",
										},
									},
								},
							},
							{
								Values: []float32{9.77, 29.64, 39.31, 32.48, 41.14, 38.38},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#00fff9",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#00a2c5",
										},
									},
								},
							},
							{
								Values: []float32{21.99, 51.04, 59.85, 52.12, 57.21, 62.28},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#3f962c",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#13761f",
										},
									},
								},
							},
							{
								Values: []float32{27.81, 54.69, 61.63, 68.50, 70.24, 81.14},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#5eab2e",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#168523",
										},
									},
								},
							},
							{
								Values: []float32{30.65, 65.00, 81.62, 63.96, 69.42, 76.28},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#ffa700",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#ff7400",
										},
									},
								},
							},
							{
								Values: []float32{34.61, 93.72, 83.89, 82.76, 102.58, 99.29},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#ffce00",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#ff8d00",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
