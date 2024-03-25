package booking

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zfirdavs/hotel-booking-test-task/pkg/model"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/storage"
)

type HotelRoomBookingService struct {
	Orders       []model.Order
	Logger       *log.Logger
	Availability storage.Storage
}

func NewHotelRoomBookingService(storage storage.Storage) *HotelRoomBookingService {
	return &HotelRoomBookingService{
		Orders:       []model.Order{},
		Logger:       log.New(os.Stdout, "", log.LstdFlags),
		Availability: storage,
	}
}

func (s *HotelRoomBookingService) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder model.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		s.handleError(w, err, "Failed to decode request body")
		return
	}

	unavailableDays := s.Availability.FindUnavailableDays(newOrder.From, newOrder.To)
	if len(unavailableDays) != 0 {
		s.handleError(w,
			errors.New("Hotel room is not available for selected dates"),
			fmt.Sprintf("Your order: %v has unvailable days: %v", newOrder, unavailableDays))
		return
	}

	s.Availability.BookRooms(newOrder.From, newOrder.To)

	s.Orders = append(s.Orders, newOrder)

	s.respondWithJSON(w, http.StatusCreated, newOrder)
	s.LogInfo("Order successfully created", newOrder)

}

func (s *HotelRoomBookingService) handleError(w http.ResponseWriter, err error, message string) {
	s.LogError(message, err)
	http.Error(w, message, http.StatusInternalServerError)
}

func (s *HotelRoomBookingService) respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *HotelRoomBookingService) LogInfo(message string, v ...interface{}) {
	s.Logger.Printf("[Info]: %s %v\n", message, v)
}

func (s *HotelRoomBookingService) LogError(message string, v ...interface{}) {
	s.Logger.Printf("[Error]: %s %v\n", message, v)
}
