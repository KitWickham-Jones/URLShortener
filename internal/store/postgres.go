package store

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Store struct{
	db *pgxpool.Pool
	rdb *redis.Client
}

func New (db *pgxpool.Pool, rdb *redis.Client) *Store{
	return &Store{
		db: db,
		rdb: rdb,
	}
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

