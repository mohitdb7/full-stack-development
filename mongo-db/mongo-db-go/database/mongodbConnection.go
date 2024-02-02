package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongoDB() (*mongo.Collection, *mongo.Client, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	collection := client.Database("mongo-golang-test").Collection("Users")

	// defer func() {
	// 	cancel()

	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Close connection is called")
	// }()

	return collection, client, ctx
}
