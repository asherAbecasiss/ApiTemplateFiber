package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiServer struct {
	Port   string
	App    *fiber.App
	Client *MongoClient
}
type MongoClient struct {
	Client *mongo.Client
	Ctx    context.Context
}

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string
}
