package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kayalova/kodix-test/config"
	"github.com/kayalova/kodix-test/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var carsCollection *mongo.Collection

func Create(c model.Car) error {
	fmt.Println(c)
	_, err := carsCollection.InsertOne(context.Background(), c)

	return err
}

func GetAll(filters interface{}) []model.Car {
	// cursor, err := carsCollection.Find(context.Background(), filters)
	// map[string]interf
	cursor, err := carsCollection.Find(context.Background(), filters)

	if err != nil {
		log.Fatal(err)
	}

	var cars []model.Car
	if err = cursor.All(context.Background(), &cars); err != nil {
		log.Fatal(err)
	}

	return cars
}

func DeleteOne(id string) {
	carsCollection.DeleteOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: id}})
}

func FindById(id string) model.Car {
	var car model.Car
	res := carsCollection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: id}})
	res.Decode(&car)
	return car
}

func init() {
	dir, _ := os.Getwd()

	err := godotenv.Load(dir + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	carsCollection = config.GetDbConnection().Database("kodix-cars").Collection("cars")
}
