package main

import (
	"errors"
	"flag"
	"net/http"
	"os"

	"github.com/zfirdavs/hotel-booking-test-task/pkg/booking"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/storage"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9000", "The port for http server")

	flag.Parse()

	availability := storage.NewAvailability()

	service := booking.NewHotelRoomBookingService(availability)
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", service.CreateOrder)

	service.LogInfo("Server listening on localhost", port)
	err := http.ListenAndServe(":"+port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		service.LogInfo("Server closed")
	} else if err != nil {
		service.LogError("Server failed", err)
		os.Exit(1)
	}
}
