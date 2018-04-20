package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

type TmProcessor struct {
	statistics map[string]opensips.Statistic
}

var tmLabelNames = []string{}
var tmMetrics = map[string]metric{
	"received_replies":   newMetric("tm", "received_replies_total", "Total number of total replies received by TM module.", tmLabelNames, prometheus.CounterValue),
	"relayed_replies":    newMetric("tm", "relayed_replies_total", "Total number of replies received and relayed by TM module.", tmLabelNames, prometheus.CounterValue),
	"local_replies":      newMetric("tm", "local_replies_total", "Total number of replies local generated by TM module.", tmLabelNames, prometheus.CounterValue),
	"UAS_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"UAC_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"2xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"3xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"4xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"5xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"6xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"inuse_transactions": newMetric("tm", "inuse_transactions", "Number of transactions existing in memory at current time.", tmLabelNames, prometheus.GaugeValue),
}

func init() {
	for metric := range tmMetrics {
		Processors[metric] = tmProcessorFunc
	}
	Processors["tm:"] = tmProcessorFunc
}

func (c TmProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range tmMetrics {
		ch <- metric.Desc
	}
}

func (p TmProcessor) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["received_replies"].Desc,
		tmMetrics["received_replies"].ValueType,
		p.statistics["received_replies"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["relayed_replies"].Desc,
		tmMetrics["relayed_replies"].ValueType,
		p.statistics["relayed_replies"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["local_replies"].Desc,
		tmMetrics["local_replies"].ValueType,
		p.statistics["local_replies"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["UAS_transactions"].Desc,
		tmMetrics["UAS_transactions"].ValueType,
		p.statistics["UAS_transactions"].Value,
		"UAS",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["UAC_transactions"].Desc,
		tmMetrics["UAC_transactions"].ValueType,
		p.statistics["UAC_transactions"].Value,
		"UAC",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["2xx_transactions"].Desc,
		tmMetrics["2xx_transactions"].ValueType,
		p.statistics["2xx_transactions"].Value,
		"2xx",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["3xx_transactions"].Desc,
		tmMetrics["3xx_transactions"].ValueType,
		p.statistics["3xx_transactions"].Value,
		"3xx",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["4xx_transactions"].Desc,
		tmMetrics["4xx_transactions"].ValueType,
		p.statistics["4xx_transactions"].Value,
		"4xx",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["5xx_transactions"].Desc,
		tmMetrics["5xx_transactions"].ValueType,
		p.statistics["5xx_transactions"].Value,
		"5xx",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["6xx_transactions"].Desc,
		tmMetrics["6xx_transactions"].ValueType,
		p.statistics["6xx_transactions"].Value,
		"6xx",
	)
	ch <- prometheus.MustNewConstMetric(
		tmMetrics["inuse_transactions"].Desc,
		tmMetrics["inuse_transactions"].ValueType,
		p.statistics["inuse_transactions"].Value,
	)
}

func tmProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &TmProcessor{
		statistics: s,
	}
}
