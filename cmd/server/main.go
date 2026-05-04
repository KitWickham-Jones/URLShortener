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
)

func main(){
	godotenv.Load()
	cfg := config.Load()
	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	st := store.New(db)
	srv := api.New(st)
	
	log.Printf("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", srv))

}
