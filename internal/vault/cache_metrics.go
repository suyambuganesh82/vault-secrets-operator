// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsNamespace     = "vso"
	subsystemClientCache = "client_cache"
)

var (

	// metricsFQNClientCacheSize for the ClientCache.
	metricsFQNClientCacheSize = prometheus.BuildFQName(
		metricsNamespace, subsystemClientCache, "size")

	// metricsFQNClientCacheLength for the ClientCache.
	metricsFQNClientCacheLength = prometheus.BuildFQName(
		metricsNamespace, subsystemClientCache, "length")

	// metricsFQNClientCacheHits for the ClientCache.
	metricsFQNClientCacheHits = prometheus.BuildFQName(
		metricsNamespace, subsystemClientCache, "hits")

	// metricsFQNClientCacheMisses for the ClientCache.
	metricsFQNClientCacheMisses = prometheus.BuildFQName(
		metricsNamespace, subsystemClientCache, "misses")

	// metricsFQNClientCacheEvictions for the ClientCache.
	metricsFQNClientCacheEvictions = prometheus.BuildFQName(
		metricsNamespace, subsystemClientCache, "evictions")
)

var _ prometheus.Collector = (*clientCacheCollector)(nil)

// clientCacheCollector provides a prometheus.Collector for ClientCache metrics.
type clientCacheCollector struct {
	cache    ClientCache
	size     float64
	sizeDesc *prometheus.Desc
	lenDesc  *prometheus.Desc
}

func (c clientCacheCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.sizeDesc
	ch <- c.lenDesc
}

func (c clientCacheCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.sizeDesc, prometheus.GaugeValue, c.size)
	ch <- prometheus.MustNewConstMetric(c.lenDesc, prometheus.GaugeValue, float64(c.cache.Len()))
}

func newClientCacheCollector(cache ClientCache, size int) prometheus.Collector {
	return &clientCacheCollector{
		cache: cache,
		size:  float64(size),
		sizeDesc: prometheus.NewDesc(
			metricsFQNClientCacheSize,
			"Size of the cache.",
			nil, nil),
		lenDesc: prometheus.NewDesc(
			metricsFQNClientCacheLength,
			"Number of Vault Clients in the cache.",
			nil, nil),
	}
}