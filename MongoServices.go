package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *MongoClient) CreateUser(ctx context.Context, user User) (User, error) {

	user.ID = primitive.NewObjectID()
	col := c.Client.Database("RMM").Collection("test")

	//ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// InsertOne() method Returns mongo.InsertOneResult
	_, insertErr := col.InsertOne(c.Ctx, user)
	if insertErr !=

		nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		// safely exit script on error
	}
	return user, nil
}
