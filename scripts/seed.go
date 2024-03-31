package main

import (
	"context"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := customTypes.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}
	rooms := []customTypes.Room{
		{
			Type:      customTypes.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      customTypes.DeluxeRoomType,
			BasePrice: 1999.9,
		},
		{
			Type:      customTypes.SeaSideRoomType,
			BasePrice: 4321.2,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelId = insertedHotel.Id
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("Bellucia", "Napoli", 100)
	seedHotel("The cozy hotel", "Roma", 250)
}

func init() {
	var err error
	// connection to the mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
	// 	log.Fatal(err)
	// }
	hotelStore := db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
