package store

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{
	db *pgxpool.Pool
}

func New (db *pgxpool.Pool) *Store{
	return &Store{db: db}
}


func (s *Store) InsertURL(ctx context.Context, slug string, longURL string ) error {
	_ , err := s.db.Exec(ctx,  
		"INSERT INTO url_map (slug, long_url) VALUES ($1, $2)", slug, longURL)
	return err
}

func (s *Store) GetURL(ctx context.Context, slug string)(string, error){
	var longURL string
	err := s.db.QueryRow(ctx,
		"SELECT long_url FROM url_map WHERE SLUG=$1", slug).Scan(&longURL)
	return longURL, err
}

