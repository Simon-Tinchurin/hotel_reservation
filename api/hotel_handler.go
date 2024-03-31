package api

import (
	"fmt"
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hs,
		roomStore:  rs,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}
	fmt.Println(qparams)

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}