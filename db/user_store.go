package db

import (
	"context"
	"fmt"
	"hotel-reservation/customTypes"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type Map map[string]any

// define separately to use across multiple interfaces
type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserByEmail(context.Context, string) (*customTypes.User, error)
	GetUserByID(context.Context, string) (*customTypes.User, error)
	GetUsers(context.Context) ([]*customTypes.User, error)
	InsertUser(context.Context, *customTypes.User) (*customTypes.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter Map, params customTypes.UpdateUserParams) error
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	dbname := os.Getenv(MongoDBEnvName)
	return &MongoUserStore{
		client:     client,
		collection: client.Database(dbname).Collection(userColl),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("--- dropping user collection ---")
	return s.collection.Drop(ctx)
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter Map, params customTypes.UpdateUserParams) error {
	oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
	if err != nil {
		return err
	}
	filter["_id"] = oid
	fmt.Println(filter)
	update := bson.M{"$set": params.ToBSON()}
	_, err = s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	// update := bson.D{{"$set", bson.D{{"email", "test"}}}}
	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// TODO: Maybe its a good idea to handle if we did not delete any user
	// maybe log it or smth
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *customTypes.User) (*customTypes.User, error) {
	res, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*customTypes.User, error) {
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*customTypes.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil

}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*customTypes.User, error) {
	var user customTypes.User
	if err := s.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*customTypes.User, error) {
	// validate the id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user customTypes.User
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
