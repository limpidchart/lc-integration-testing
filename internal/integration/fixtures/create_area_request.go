package fixtures

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-integration-testing/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func CreateAreaRequest() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Title: "Single Area Chart",
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
				Kind:       render.ChartScale_BAND,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 1130},
				Domain: &render.ChartScale_DomainCategories{
					DomainCategories: &render.DomainCategories{
						Categories: []string{"a1", "a2", "a3", "a4", "a5", "a6"},
					},
				},
				NoBoundariesOffset: true,
				InnerPadding:       &wrapperspb.FloatValue{Value: 0},
				OuterPadding:       &wrapperspb.FloatValue{Value: 0},
			},
			AxisBottomLabel: "X Values",
			AxisLeft: &render.ChartScale{
				Kind:       render.ChartScale_LINEAR,
				RangeStart: &wrapperspb.Int32Value{Value: 0},
				RangeEnd:   &wrapperspb.Int32Value{Value: 610},
				Domain: &render.ChartScale_DomainNumeric{
					DomainNumeric: &render.DomainNumeric{
						Start: 0,
						End:   200,
					},
				},
			},
			AxisLeftLabel: "Y Values",
		},
		Views: []*render.ChartView{
			{
				Kind: render.ChartView_AREA,
				Values: &render.ChartView_ScalarValues{
					ScalarValues: &render.ChartViewScalarValues{
						Values: []float32{12, 100, 120, 180, 40, 8},
					},
				},
				Colors: &render.ChartViewColors{
					Fill: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#ffa700",
						},
					},
					Stroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#ff8d00",
						},
					},
					PointFill: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#ff7400",
						},
					},
					PointStroke: &render.ChartElementColor{
						ColorValue: &render.ChartElementColor_ColorHex{
							ColorHex: "#ff7400",
						},
					},
				},
				PointType:          render.ChartView_SQUARE,
				PointLabelPosition: render.ChartView_TOP_RIGHT,
			},
		},
	}
}
