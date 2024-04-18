package db

import (
	"context"
	"hotel-reservation/customTypes"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *customTypes.Room) (*customTypes.Room, error)
	GetRooms(context.Context, bson.M) ([]*customTypes.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	dbname := os.Getenv(MongoDBEnvName)
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*customTypes.Room, error) {
	resp, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*customTypes.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *customTypes.Room) (*customTypes.Room, error) {
	resp, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = resp.InsertedID.(primitive.ObjectID)
	// update the hotel with this room id
	filter := Map{"_id": room.HotelId}
	update := Map{"$push": bson.M{"rooms": room.Id}}
	if err := s.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}
