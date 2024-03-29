package main

import (
	"context"
	"fmt"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	// connection to the mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := customTypes.Hotel{
		Name:     "Bellucia",
		Location: "Narnia",
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
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
}
