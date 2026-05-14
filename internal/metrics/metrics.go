package metrics

import "github.com/prometheus/client_golang/prometheus"



type Metrics struct{
	redirectsTotal prometheus.Counter
	cacheHits prometheus.Counter
	cacheMisses prometheus.Counter
	requestDuration prometheus.Histogram
}

func New() *Metrics{
	m := &Metrics{
		redirectsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "redirects_total",
			Help: "Total number of redirect requests",
		}),
		cacheHits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of cache hits",
		}),
		cacheMisses: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of cache misses",
		}),
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "request_duration_seconds",
			Help: "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		}),
	}
	prometheus.MustRegister(
			m.redirectsTotal,
			m.cacheHits,
			m.cacheMisses,
			m.requestDuration,
		)
	return m
}