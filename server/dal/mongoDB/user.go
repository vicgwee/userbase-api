package mongoDB

import (
	"context"
	"fmt"

	"userbase-api/server/dal"
	ue "userbase-api/server/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "userDB"
	collectionName = "userData"
)

var userCollection *mongo.Collection

func InitUserCollection(ctx context.Context) error {
	db := client.Database("userDB")
	if db == nil {
		return ue.NewMongoDbError("InitUserCollection failed to connect to db=%s", dbName)
	}
	userCollection = db.Collection("userData")
	if userCollection == nil {
		return ue.NewMongoDbError("InitUserCollection failed to connect to collection=%s", collectionName)
	}
	docCount, err := userCollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return ue.NewMongoDbError("InitUserCollection failed to count documents, err=%v", err)
	}
	fmt.Printf("InitUserCollection connected to %s:%s, docCount=%v\n", dbName, collectionName, docCount)
	return nil
}

// The repository design pattern is used to decouple the database interface and implementation
// It also makes unit testing easy!
type UserRepository interface {
	CreateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error)
	UpdateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error)
	GetUser(ctx context.Context, id string) (*dal.UserDocument, error)
	GetUsers(ctx context.Context) ([]*dal.UserDocument, error)
	DeleteUser(ctx context.Context, id string) error
}

type RepositoryImpl struct{}

func NewUserRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error) {
	_, err := userCollection.InsertOne(ctx, doc)
	if err != nil {
		return nil, ue.NewInternalError("CreateUser failed, doc=%v, err=%v", doc, err)
	}
	return r.GetUser(ctx, *doc.Id)
}

func (r *RepositoryImpl) UpdateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error) {
	filter := bson.D{{Key: "id", Value: doc.Id}}
	update := bson.D{{Key: "$set", Value: doc}}
	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, ue.NewInternalError("UpdateUser failed, doc=%v, err=%v", doc, err)
	}
	return r.GetUser(ctx, *doc.Id)
}

func (r *RepositoryImpl) GetUser(ctx context.Context, id string) (*dal.UserDocument, error) {
	filter := bson.D{{Key: "id", Value: id}}
	docs, err := userCollection.Find(ctx, filter)
	if err != nil || docs == nil {
		return nil, ue.NewUserNotFoundError("GetUser failed to find user, id=%s, err=%v", id, err)
	}

	var users []*dal.UserDocument
	err = docs.All(ctx, &users)
	if err != nil {
		return nil, ue.NewInternalError("GetUser failed to convert docs to users, id=%s, err=%v", id, err)
	}

	if len(users) != 1 {
		return nil, ue.NewUserDuplicateError("GetUser returned count != 1, id=%s, users=%v", id, users)
	}

	return users[0], nil
}

func (r *RepositoryImpl) GetUsers(ctx context.Context) ([]*dal.UserDocument, error) {
	filter := bson.D{}
	docs, err := userCollection.Find(ctx, filter)
	if err != nil {
		return nil, ue.NewUserNotFoundError("GetUsers failed to find users, err=%v", err)
	}

	var users []*dal.UserDocument
	err = docs.All(ctx, &users)
	if err != nil {
		return nil, ue.NewInternalError("GetUsers failed to convert docs to users, err=%v", err)
	}

	return users, nil
}

func (r *RepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return ue.NewInternalError("DeleteUser failed, id=%v, err=%v", id, err)
	}
	return nil
}
