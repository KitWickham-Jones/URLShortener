package main

import (
	"context"
	"log"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kitwj/urlshortener/internal/api"
	"github.com/kitwj/urlshortener/internal/config"
	"github.com/kitwj/urlshortener/internal/store"
	"github.com/redis/go-redis/v9"
)

func main(){
	godotenv.Load()
	cfg := config.Load()
	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil{
		log.Fatal("Could not connect to database.", err)
	}
	log.Println("Database connected")

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisURL,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil{
		log.Fatal("Could not connect to redis", err)
	}


	st := store.New(db, rdb)
	srv := api.New(st, cfg)
	
	log.Printf("server started at %s", cfg.BaseURL)
	log.Fatal(http.ListenAndServe(cfg.Port, srv))

}
