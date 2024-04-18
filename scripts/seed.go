package main

import (
	"context"
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/db"
	"hotel-reservation/db/fixtures"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	// connection to the mongodb
	fmt.Println(mongoDBName)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
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

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("hotel_name_%d", i)
		location := fmt.Sprintf("location_%d", i)
		fixtures.AddHotel(&store, name, location, rand.Intn(5)+1, nil)
	}
}
