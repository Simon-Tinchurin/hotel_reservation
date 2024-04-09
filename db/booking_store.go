package db

import (
	"context"
	"hotel-reservation/customTypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	InsertBooking(context.Context, *customTypes.Booking) (*customTypes.Booking, error)
	GetBookings(context.Context, bson.M) ([]*customTypes.Booking, error)
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	BookingStore
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("bookings"),
	}
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *customTypes.Booking) (*customTypes.Booking, error) {
	resp, err := s.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.Id = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*customTypes.Booking, error) {
	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*customTypes.Booking
	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}
