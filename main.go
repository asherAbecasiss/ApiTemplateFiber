package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *MongoClient {
	mongoOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.Background(), mongoOptions)
	ct := context.Background()

	if err != nil {
		panic(err)
	}

	return &MongoClient{
		Client: c,
		Ctx:    ct,
	}
}

func GetApiServer(data *MongoClient) *ApiServer {
	app := fiber.New()
	return &ApiServer{
		Port:   ":8080",
		App:    app,
		Client: data,
	}

}

func main() {
	cl := NewMongoClient()
	app := GetApiServer(cl)
	app.Run()

}
