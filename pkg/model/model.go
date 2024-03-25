package model

import (
	"fmt"
	"time"
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o Order) String() string {
	return fmt.Sprintf("'Email: %s From: %v To: %v'", o.UserEmail, o.From.Format("2006-01-02"), o.To.Format("2006-01-02"))
}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

type RoomAvailabilityV2 struct {
	HotelID string `json:"hotel_id"`
	RoomID  string `json:"room_id"`
	Quota   int    `json:"quota"`
}
