package domain

import (
	"context"
	"fmt"
	"log"
	"mongo-db-go/database"
	"mongo-db-go/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	CreateUser(dto.User) (*dto.User, error)
	FindUser(string) (*dto.User, error)
	DeleteUser(string) (*string, error)
	FindUsers() ([]dto.User, error)
}

type UserRepoMongo struct {
}

func (urm UserRepoMongo) CreateUser(user dto.User) (*dto.User, error) {

	collection, client, context, cancel := database.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	user.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context, user)
	if err != nil {
		fmt.Printf("Error in writing object %s", err)
		return nil, err
	}

	fmt.Printf("Inserted user %s", result.InsertedID)

	return &user, nil
}

func (urm UserRepoMongo) FindUser(id string) (*dto.User, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := database.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	u := dto.User{}
	filter := bson.D{{Key: "_id", Value: oid}}

	err = collection.FindOne(context, filter).Decode(&u)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document found")
		return nil, err
	} else if err != nil {
		fmt.Printf("Error in mongo %v", err)
		return nil, err
	}

	return &u, nil
}

func (urm UserRepoMongo) FindUsers() ([]dto.User, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := database.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	u := make([]dto.User, 0, 10)
	filter := bson.D{}

	cursor, err := collection.Find(context, filter)
	defer cursor.Close(context)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No document found")
		return nil, err
	} else if err != nil {
		fmt.Printf("Error in mongo %v", err)
		return nil, err
	}

	for cursor.Next(context) {
		var result dto.User
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		u = append(u, result)
	}

	return u, nil
}

func (urm UserRepoMongo) DeleteUser(id string) (*string, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := database.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	filter := bson.D{{Key: "_id", Value: oid}}

	result, err := collection.DeleteOne(context, filter)
	if err != nil {
		fmt.Printf("Error with delete %v", err)
	}

	resultStr := fmt.Sprintf("Results deleted %v, id of the item %v", result.DeletedCount, oid)
	return &resultStr, nil
}

func closeConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()

		if err := client.Disconnect(context); err != nil {
			panic(err)
		}

		fmt.Println("Close connection is called")
	}()
}

func NewUserRepoMongo() IUserRepo {
	return UserRepoMongo{}
}
