package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type BotMetrics struct {
	unaryRequestCounter prometheus.Counter
}

func NewBotMetrics() *BotMetrics {
	return &BotMetrics{
		unaryRequestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "bot",
			Name:      "unary_requests_counter",
			Help:      "total count",
		}),
	}
}

func (m BotMetrics) UnaryRequestInc() {
	m.unaryRequestCounter.Inc()
}
