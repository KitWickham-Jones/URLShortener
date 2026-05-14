package api

import (
	"log"
	"net/http"
	"time"
	"github.com/kitwj/urlshortener/internal/config"
	"github.com/kitwj/urlshortener/internal/metrics"
	"github.com/kitwj/urlshortener/internal/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct{
	store *store.Store
	router *http.ServeMux
	config *config.Config
	metrics *metrics.Metrics
}

func New (st *store.Store, cfg *config.Config, met * metrics.Metrics) *Server{
	s := &Server{
		store: st,
		config: cfg,
	}
	s.router = http.NewServeMux()
	s.router.HandleFunc("POST /shorten", s.handleShorten)
	s.router.HandleFunc("GET /{slug}", s.handleRedirect)
	s.router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})
	s.router.Handle("GET /metrics", promhttp.Handler())
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("%s %s", r.Method, r.URL.Path)
	s.router.ServeHTTP(w, r)
	s.metrics.RequestDuration.Observe(time.Since(start).Seconds())
}