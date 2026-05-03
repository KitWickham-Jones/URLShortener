package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/kitwj/urlshortener/internal/config"
)

func main(){
	godotenv.Load()
	cfg := config.Load()
	db, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close(context.Background())

}

// package main

// import(
// 	"fmt"
// 	"github.com/kitwj/urlshortener/internal/api"
// )

// func main(){
// 	fmt.Println(api.GenerateSlug())
// }