package api

import (
	"log"
	"net/http"
	"github.com/kitwj/urlshortener/internal/config"
	"github.com/kitwj/urlshortener/internal/store"
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
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	s.router.ServeHTTP(w, r)
}