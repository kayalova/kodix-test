package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kayalova/kodix-test/config"
)

func init() {
	dir, _ := os.Getwd()
	err := godotenv.Load(dir + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	// port, _ := config.GetEnv("APP_PORT")
	config.GetEnv("APP_PORT")
	config.GetDbConnection()
}
