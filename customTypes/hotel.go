package customTypes

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	_              RoomType = iota // to start cont from 1
	SingleRoomType                 // 1
	DoubleRoomType                 // 2
	SeaSideRoomType
	DeluxeRoomType
)

type Hotel struct {
	Id       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type RoomType int

type Room struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// small, normal, kingsize
	Size    string             `bson:"size" json:"size"`
	Seaside bool               `bson:"seaside" json:"seaside"`
	Price   float64            `bson:"price" json:"price"`
	HotelId primitive.ObjectID `bson:"hotelId" json:"hotelId"`
}
