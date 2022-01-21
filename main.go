package main

import (
	"log"
	db "stock-controller/internal/db/mongo"
	"stock-controller/pkg/routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Mongo.Init()
	e := routes.Handler()
	e.Logger.Fatal(e.Start(":8005"))
}
