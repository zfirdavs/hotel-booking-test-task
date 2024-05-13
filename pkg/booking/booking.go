package booking

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zfirdavs/hotel-booking-test-task/pkg/logger"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/model"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/storage"
)

type HotelRoomBookingService struct {
	Logger       logger.Log
	Availability storage.Storage
}

func NewHotelRoomBookingService(storage storage.Storage, log logger.Log) *HotelRoomBookingService {
	return &HotelRoomBookingService{
		Logger:       log,
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
			fmt.Errorf("hotel room is not available for selected dates: %v", unavailableDays),
			fmt.Sprintf("Order: %v has unvailable days: %v", newOrder, unavailableDays),
		)
		return
	}

	s.Availability.BookRooms(newOrder.From, newOrder.To)

	s.respondWithJSON(w, http.StatusCreated, newOrder)
	s.LogInfo("Order successfully created", "order_info", newOrder)

}

func (s *HotelRoomBookingService) handleError(w http.ResponseWriter, err error, message string) {
	s.LogError(message)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (s *HotelRoomBookingService) respondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *HotelRoomBookingService) LogInfo(message string, v ...any) {
	s.Logger.Info(message, v...)
}

func (s *HotelRoomBookingService) LogError(message string, v ...any) {
	s.Logger.Error(message, v...)
}
