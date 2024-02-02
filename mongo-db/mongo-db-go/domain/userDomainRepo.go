package domain

import (
	"context"
	"fmt"
	"mongo-db-go/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	CreateUser(dto.User) (*dto.User, error)
}

type UserRepoMongo struct {
	collection *mongo.Collection
	clinet     *mongo.Client
	context    context.Context
}

func (urm UserRepoMongo) CreateUser(user dto.User) (*dto.User, error) {
	user.Id = primitive.NewObjectID()
	result, err := urm.collection.InsertOne(urm.context, user)
	if err != nil {
		fmt.Printf("Error in writing object %s", err)
		return nil, err
	}

	fmt.Printf("Inserted user %s", result.InsertedID)

	return &user, nil

}

func NewUserRepoMongo(collection *mongo.Collection, client *mongo.Client, context context.Context) IUserRepo {
	return UserRepoMongo{
		collection: collection,
		clinet:     client,
		context:    context,
	}
}
