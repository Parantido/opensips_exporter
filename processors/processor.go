package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

type metric struct {
	Desc      *prometheus.Desc
	ValueType prometheus.ValueType
}

type processor func(map[string]opensips.Statistic) prometheus.Collector

var Processors = make(map[string]processor)

const namespace = "opensips"

func newMetric(subsystem string, name string, help string, variableLabels []string, t prometheus.ValueType) metric {
	return metric{
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, name),
			help, variableLabels, nil,
		), t}
}
