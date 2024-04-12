package main

import (
	"context"
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/db"
	"hotel-reservation/db/fixtures"
	"log"
	"time"

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
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)

	store := db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(&store, "james", "foo", false)
	admin := fixtures.AddUser(&store, "admin", "admin", true)
	fmt.Printf("James -> %s\n", api.CreateTokenFromUser(user))
	fmt.Printf("Admin -> %s\n", api.CreateTokenFromUser(admin))

	hotel := fixtures.AddHotel(&store, "some hotel", "bermuda", 5, nil)
	// fmt.Println(hotel)
	room := fixtures.AddRoom(&store, "large", true, 543.22, hotel.Id)
	// fmt.Println(room)
	booking := fixtures.AddBooking(&store, user.Id, room.Id, time.Now(), time.Now().AddDate(0, 0, 3))
	fmt.Printf("Booking ID -> %s\n", booking.Id)

}
