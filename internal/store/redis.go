package store

import (
	"context"
	"time"
)

func (s *Store) GetCachedURL(ctx context.Context, slug string) (string,error){
	return s.rdb.Get(ctx, slug).Result()
}

func (s *Store) CacheURL( ctx context.Context, slug string, longURL string) error { 
	return s.rdb.Set(ctx, slug, longURL, 24*time.Hour).Err()
}