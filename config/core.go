package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Execute() {
	loadEnvironment()
	// mongo.ConnectionToMongoDB()

	fmt.Println("[Config] Loaded Succesful")
}

func loadEnvironment() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error to load .env file", err)
	}
}
