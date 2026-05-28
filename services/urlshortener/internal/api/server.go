package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/kitwj/urlshortener/internal/config"
	"github.com/kitwj/urlshortener/internal/metrics"
	"github.com/kitwj/urlshortener/internal/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
)

type Server struct{
	store *store.Store
	router *http.ServeMux
	config *config.Config
	metrics *metrics.Metrics
	kafka *kafka.Writer
}

func New (st *store.Store, cfg *config.Config, met *metrics.Metrics, kfk *kafka.Writer) *Server{
	s := &Server{
		store: st,
		config: cfg,
		metrics: met,
		kafka: kfk,
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

func (s *Server) publishClickEvent(slug string){
	log.Printf("publishing click event for %s", slug)
	payload, _ := json.Marshal(map[string]string{
		"slug": slug,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})

	if err := s.kafka.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	}); err != nil {
		log.Printf("kafka write error: %v", err)
	}
}