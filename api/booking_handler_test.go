package api

import (
	"fmt"
	"hotel-reservation/db/fixtures"
	"testing"
	"time"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "james", "foo", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 5.234, hotel.Id)
	from := time.Now()
	till := time.Now().AddDate(0, 0, 10)
	booking := fixtures.AddBooking(db.Store, user.Id, room.Id, from, till)
	fmt.Println(booking)
}
