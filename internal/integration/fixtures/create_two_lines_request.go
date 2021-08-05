package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateTwoLinesRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 1000},
			Height: &wrapperspb.Int32Value{Value: 800},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 30},
			MarginBottom: &wrapperspb.Int32Value{Value: 40},
			MarginLeft:   &wrapperspb.Int32Value{Value: 40},
			MarginRight:  &wrapperspb.Int32Value{Value: 20},
		},
		Axes: &render.ChartAxes{
			AxisBottom: &render.ChartScale{
				Kind: render.ChartScale_BAND,
				Domain: &render.ChartScale_DomainCategories{
					DomainCategories: &render.DomainCategories{
						Categories: []string{
							"A",
							"B",
							"C",
							"D",
							"E",
							"F",
							"G",
							"H",
						},
					},
				},
			},
			AxisLeft: &render.ChartScale{
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
				Kind: render.ChartView_LINE,
				Values: &render.ChartView_ScalarValues{
					ScalarValues: &render.ChartViewScalarValues{
						Values: []float32{20, 70, 130, 180, 20, 77, 140, 190},
					},
				},
				PointLabelPosition: render.ChartView_BOTTOM_LEFT,
				Colors: &render.ChartViewColors{
					Stroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#0e3569",
						},
					},
				},
				PointVisible: &wrapperspb.BoolValue{Value: false},
			},
			{
				Kind: render.ChartView_LINE,
				Values: &render.ChartView_ScalarValues{
					ScalarValues: &render.ChartViewScalarValues{
						Values: []float32{54, 40, 50, 77, 91, 53, 11, 3},
					},
				},
				PointLabelPosition: render.ChartView_TOP_RIGHT,
				Colors: &render.ChartViewColors{
					Stroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#5095e5",
						},
					},
				},
				PointVisible: &wrapperspb.BoolValue{Value: false},
			},
		},
	}
}
