package main

import (
	"github.com/joho/godotenv"
	"log"
	"github.com/jaebradley/jaeurls/store"
	"github.com/jaebradley/jaeurls/router"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	session := store.CreateStore()
	defer session.Close()

	router.StartRouter(session)
}
