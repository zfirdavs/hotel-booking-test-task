package storage

import (
	"sync"
	"time"

	"github.com/zfirdavs/hotel-booking-test-task/pkg/datetime"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/model"
)

type Availability struct {
	value map[time.Time]*model.RoomAvailabilityV2
	mu    sync.Mutex
}

func NewAvailability() *Availability {
	return &Availability{
		value: map[time.Time]*model.RoomAvailabilityV2{
			datetime.NewDate(2024, 1, 1): {
				HotelID: "reddison",
				RoomID:  "lux",
				Quota:   1,
			},
			datetime.NewDate(2024, 1, 2): {
				HotelID: "reddison",
				RoomID:  "lux",
				Quota:   1,
			},
			datetime.NewDate(2024, 1, 3): {
				HotelID: "reddison",
				RoomID:  "lux",
				Quota:   1,
			},
			datetime.NewDate(2024, 1, 5): {
				HotelID: "reddison",
				RoomID:  "lux",
				Quota:   0,
			},
		},
	}
}

func (a *Availability) FindUnavailableDays(from, to time.Time) []string {
	daysToBook := datetime.DaysBetween(from, to)
	unavailableDays := make([]string, 0)

	for _, dayToBook := range daysToBook {
		result, ok := a.value[dayToBook]
		if !ok || result.Quota < 1 {
			unavailableDays = append(unavailableDays, dayToBook.Format("2006-01-02"))
		}
	}

	return unavailableDays
}

func (a *Availability) BookRooms(from, to time.Time) {
	daysToBook := datetime.DaysBetween(from, to)

	for _, dayToBook := range daysToBook {
		res, ok := a.value[dayToBook]
		if ok && res.Quota > 0 {
			a.mu.Lock()
			res.Quota -= 1
			a.mu.Unlock()
		}
	}
}
