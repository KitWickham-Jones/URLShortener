package api

import (
	"github.com/kitwj/urlshortener/internal/store"
)

type Server struct{
	store *store.Store
}

func New (st *store.Store) *Server{
	return &Server{store: st}
}
