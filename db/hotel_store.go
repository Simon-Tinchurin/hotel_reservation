package db

import (
	"context"
	"hotel-reservation/customTypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *customTypes.Hotel) (*customTypes.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(dbname).Collection("hotels"),
	}
}

func (s *MongoHotelStore) Update(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *customTypes.Hotel) (*customTypes.Hotel, error) {
	resp, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.Id = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
