// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	vmtotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "vm_autoscaler_total",
			Help: "弹性扩容节点数",
		},
		[]string{"provider"},
	)
)

func init() {
	metrics.Registry.MustRegister(vmtotal)
}
