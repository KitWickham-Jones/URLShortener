package store

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)

func (s *Store) GetCachedURL(ctx context.Context, slug string) (string,error){
	url, err := s.rdb.Get(ctx, slug).Result()
	if err == redis.Nil{
		s.metrics.CacheMisses.Inc()
		return "", err
	} else if err != nil{
		return "", err
	}
	s.metrics.CacheHits.Inc()
	return url, nil
}

func (s *Store) CacheURL( ctx context.Context, slug string, longURL string) error { 
	return s.rdb.Set(ctx, slug, longURL, 24*time.Hour).Err()
}