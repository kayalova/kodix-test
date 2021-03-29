package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetEnv returns env variable and error
func GetEnv(key string) (string, error) {

	if val, isExists := os.LookupEnv(key); isExists {
		return val, nil
	}

	return "", fmt.Errorf("Can not find %s env variable", key)
}

// GetDbConnection returns mongo connection
func GetDbConnection() *mongo.Client {
	mongoURI, err := GetEnv("MONGO_URI")
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
