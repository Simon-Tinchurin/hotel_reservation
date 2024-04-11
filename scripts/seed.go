package main

import (
	"context"
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	userStore    db.UserStore
	bookingStore db.BookingStore
	ctx          = context.Background()
)

func seedUser(isAdmin bool, fname, lname, email, password string) *customTypes.User {
	user, err := customTypes.NewUserFromParams(customTypes.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = isAdmin
	insertedUser, err := userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
	return insertedUser
}

func seedHotel(name, location string, rating int) *customTypes.Hotel {
	hotel := customTypes.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	return insertedHotel
}

func seedRoom(size string, ss bool, price float64, hotelId primitive.ObjectID) *customTypes.Room {
	room := &customTypes.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelId: hotelId,
	}
	insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func seedBooking(roomId, userId primitive.ObjectID, from, till time.Time) {
	booking := &customTypes.Booking{
		UserId:   userId,
		RoomId:   roomId,
		FromDate: from,
		TillDate: till,
	}
	resp, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("booking -> %s\n", resp.Id)
}

func init() {
	var err error
	// connection to the mongodb
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
	// 	log.Fatal(err)
	// }
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)
}

func main() {
	seedUser(true, "admin", "admin", "admin@admin.com", "admin1234")
	james := seedUser(false, "James", "D", "james@d.com", "james1234")
	// seedHotel("Bellucia", "Napoli", 100)
	hotel := seedHotel("The cozy hotel", "Roma", 250)
	seedRoom("small", false, 34.99, hotel.Id)
	seedRoom("medium", false, 345.99, hotel.Id)
	room := seedRoom("large", true, 400.99, hotel.Id)
	seedBooking(room.Id, james.Id, time.Now(), time.Now().AddDate(0, 0, 2))

}
