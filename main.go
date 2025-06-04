package main

import (
	"log"
	"spinzoneid/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Cannot load .env file", err)
	}

	app.InitializeDB()
	r := app.InitializeRoute()
	r.Run()
}
