package fixtures

import (
	"context"
	"fmt"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *db.Store, fname, lname string, admin bool) *customTypes.User {

	user, err := customTypes.NewUserFromParams(customTypes.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Email:     fmt.Sprintf("%s@%s.com", fname, lname),
		Password:  fmt.Sprintf("%s_%s", fname, lname),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddHotel(store *db.Store, name, loc string, rating int, rooms []primitive.ObjectID) *customTypes.Hotel {
	var roomIds = rooms
	if rooms == nil {
		roomIds = []primitive.ObjectID{}
	}
	hotel := customTypes.Hotel{
		Name:     name,
		Location: loc,
		Rooms:    roomIds,
		Rating:   rating,
	}

	insertedHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddRoom(store *db.Store, size string, ss bool, price float64, hid primitive.ObjectID) *customTypes.Room {
	room := &customTypes.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelId: hid,
	}
	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddBooking(store *db.Store, userid, rid primitive.ObjectID, from, till time.Time) *customTypes.Booking {
	booking := &customTypes.Booking{
		UserId:   userid,
		RoomId:   rid,
		FromDate: from,
		TillDate: till,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
