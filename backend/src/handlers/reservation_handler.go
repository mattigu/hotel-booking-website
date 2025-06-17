package handlers

import (
	"bd2_projekt/database"
	"bd2_projekt/services"
	"bd2_projekt/schemas"
	"encoding/json"
	"net/http"
)

type ReservationHandler struct {
	service *services.ReservationService
}

func NewReservationHandler(db *database.Database) *ReservationHandler {
	reservationService := services.NewReservationService(db)
	return &ReservationHandler{service: reservationService}
}

func (reservationHandler *ReservationHandler) ReserveRoom (responseWriter http.ResponseWriter, req *http.Request) error {
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(responseWriter, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return nil
	}

	var reservationData schemas.Reservation;
	err := json.NewDecoder(req.Body).Decode(&reservationData);

	if err != nil {
		http.Error(responseWriter, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return err
	}

	reservationHandler.service.ReserveRoom(&reservationData)
	
	responseWriter.WriteHeader(http.StatusOK)
	return err
}

func (handler *ReservationHandler) GetReservationsFor (responseWriter http.ResponseWriter, req *http.Request) error {
	var userData schemas.UserData;

	userData.Name = req.URL.Query().Get("name")
	userData.Surname = req.URL.Query().Get("surname")
	userData.PhoneNumber = req.URL.Query().Get("phone_number")

	reservations := handler.service.GetReservationsForUser(&userData)
	
	return json.NewEncoder(responseWriter).Encode(reservations)
}