package storage

import "time"

type Storage interface {
	FindUnavailableDays(from, to time.Time) []string
	BookRooms(from, to time.Time)
}
