package store

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kitwj/urlshortener/internal/metrics"
	"github.com/redis/go-redis/v9"
)


type Store struct{
	db *pgxpool.Pool
	rdb *redis.Client
	metrics *metrics.Metrics
}

func New (db *pgxpool.Pool, rdb *redis.Client, met *metrics.Metrics) *Store{
	return &Store{
		db: db,
		rdb: rdb,
	}
}
