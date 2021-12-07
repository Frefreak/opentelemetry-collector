package otlpreceiver

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	http404RequestCount = stats.Int64("otlp_receiver_http_404_request_count", "Number of 404 requests", stats.UnitDimensionless)
)

func MetricViews() []*view.View {
	return []*view.View{
		{
			Name:        http404RequestCount.Name(),
			Description: http404RequestCount.Description(),
			Measure:     http404RequestCount,
			TagKeys:     []tag.Key{},
			Aggregation: view.Sum(),
		},
	}
}
