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
	GetHotelById(context.Context, primitive.ObjectID) (*customTypes.Hotel, error)
	GetHotels(context.Context, bson.M) ([]*customTypes.Hotel, error)
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("hotels"),
	}
}

func (s *MongoHotelStore) GetHotelById(ctx context.Context, id primitive.ObjectID) (*customTypes.Hotel, error) {
	var hotel customTypes.Hotel
	if err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*customTypes.Hotel, error) {
	resp, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*customTypes.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
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
