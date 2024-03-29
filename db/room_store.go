package db

import (
	"context"
	"hotel-reservation/customTypes"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *customTypes.Room) (*customTypes.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection("rooms"),
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *customTypes.Room) (*customTypes.Room, error) {
	resp, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = resp.InsertedID.(primitive.ObjectID)
	// update the hotel with this room id
	// filter := bson.M{"_id": room.HotelId}
	// update := bson.M{"$push", bson.M{"rooms": room.Id}}

	return room, nil
}
