package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB() error {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(
		context.Background(),
		clientOptions,
	)

	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return err
	}

	Client = client

	fmt.Println("Mongo connected successfully")

	return nil

}
