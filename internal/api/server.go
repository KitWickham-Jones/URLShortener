package api

import (
	"log"
	"net/http"
	"github.com/kitwj/urlshortener/internal/store"
)

type Server struct{
	store *store.Store
	router *http.ServeMux
}

func New (st *store.Store) *Server{
	s := &Server{store: st}
	s.router = http.NewServeMux()
	s.router.HandleFunc("/shorten", s.handleShorten)
	return s
}


func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	s.router.ServeHTTP(w, r)
}