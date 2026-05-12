package api

import (
	"log"
	"net/http"
	"github.com/kitwj/urlshortener/internal/config"
	"github.com/kitwj/urlshortener/internal/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct{
	store *store.Store
	router *http.ServeMux
	config *config.Config
}

func New (st *store.Store, cfg *config.Config) *Server{
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
	log.Printf("%s %s", r.Method, r.URL.Path)
	s.router.ServeHTTP(w, r)
}