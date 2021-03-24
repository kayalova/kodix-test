package config

import (
	"fmt"
	"log"
	"os"

	"github.com/goonode/mogo"
)

// GetEnv returns env variable and error
func GetEnv(key string) (string, error) {

	if val, isExists := os.LookupEnv(key); isExists {
		return val, nil
	}

	return "", fmt.Errorf("Can not find %s env variable", key)
}

// GetDbConnection returns mongo connection
func GetDbConnection() *mogo.Connection {
	dbUri, err := GetEnv("MONGODB_URI")

	if err != nil {
		log.Fatal("Can not connect to db: ", err)
	}

	const database = "kodix-cars"
	config := &mogo.Config{
		ConnectionString: dbUri,
		Database:         database,
	}

	connection, err := mogo.Connect(config)

	if err != nil {
		log.Fatal("Can not connect to db: ", err)
	}

	return connection
}
