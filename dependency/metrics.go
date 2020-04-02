// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package dependency

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsNamespace = "juju"
	metricsSubsystem = "worker"

	groupLabel = "group"
	nameLabel  = "name"
)

var (
	workerStartLabelNames = []string{
		groupLabel,
		nameLabel,
	}
)

// metrics holds the prometheus metrics that are exposed for
// the dependency engine.
type metrics struct {
	WorkerStarts *prometheus.CounterVec
}

func createMetrics() *metrics {
	return &metrics{
		WorkerStarts: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "worker_starts",
			Help:      "The humber of times a worker has started",
		}, workerStartLabelNames),
	}
}

// Describe is part of the prometheus.Collector interface.
func (e *Engine) Describe(ch chan<- *prometheus.Desc) {
	e.metrics.WorkerStarts.Describe(ch)
}

// Collect is part of the prometheus.Collector interface.
func (e *Engine) Collect(ch chan<- prometheus.Metric) {
	e.metrics.WorkerStarts.Collect(ch)
}
