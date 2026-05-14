package metrics

import "github.com/prometheus/client_golang/prometheus"



type Metrics struct{
	RedirectsTotal prometheus.Counter
	CacheHits prometheus.Counter
	CacheMisses prometheus.Counter
	RequestDuration prometheus.Histogram
}

func New() *Metrics{
	m := &Metrics{
		RedirectsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "redirects_total",
			Help: "Total number of redirect requests",
		}),
		CacheHits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of cache hits",
		}),
		CacheMisses: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of cache misses",
		}),
		RequestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "request_duration_seconds",
			Help: "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		}),
	}
	prometheus.MustRegister(
			m.RedirectsTotal,
			m.CacheHits,
			m.CacheMisses,
			m.RequestDuration,
		)
	return m
}