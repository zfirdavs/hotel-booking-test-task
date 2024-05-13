package main

import (
	"errors"
	"flag"
	"net/http"
	"os"

	"github.com/zfirdavs/hotel-booking-test-task/pkg/booking"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/logger"
	"github.com/zfirdavs/hotel-booking-test-task/pkg/storage"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9000", "The port for http server")

	flag.Parse()

	availability := storage.NewAvailability()

	jsonLogger := logger.NewJSONLogger(os.Stdout)

	service := booking.NewHotelRoomBookingService(availability, jsonLogger)
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", service.CreateOrder)

	jsonLogger.Info("Server listening on localhost", "port", port)

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		jsonLogger.Info("Server closed")
	} else if err != nil {
		jsonLogger.Error("Server failed", err)
		os.Exit(1)
	}
}
