package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateVerticalBarRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Title: "Vertical Bar Chart",
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
				Kind: render.ChartScale_BAND,
				Domain: &render.ChartScale_DomainCategories{
					DomainCategories: &render.DomainCategories{
						Categories: []string{"A", "B", "C", "D", "E"},
					},
				},
			},
			AxisBottomLabel: "Categories",
			AxisLeft: &render.ChartScale{
				Kind: render.ChartScale_LINEAR,
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   100,
					},
				},
			},
			AxisLeftLabel: "Values",
		},
		Views: []*render.ChartView{
			{
				Kind: render.ChartView_VERTICAL_BAR,
				Values: &render.ChartView_BarsValues{
					BarsValues: &render.ChartViewBarsValues{
						BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
							{
								Values: []float32{92, 12, 34.8, 24, 2.5},
								Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
									Fill: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#36896e",
										},
									},
									Stroke: &render.ChartElementColor{
										ColorValue: &render.ChartElementColor_ColorHex{
											ColorHex: "#0c513b",
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
