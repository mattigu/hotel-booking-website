package services

import (
	"bd2_projekt/database"
	"bd2_projekt/repositories"
	"bd2_projekt/schemas"
)

type ReservationService struct {
	repository *repositories.ReservationRepository
}

func NewReservationService(db *database.Database) *ReservationService {
	reservationRepository := &repositories.ReservationRepository{Db: db}
	return &ReservationService{repository: reservationRepository}
}

func (reservationService *ReservationService) ReserveRoom(reservationData *schemas.Reservation) error {
	return reservationService.repository.ReserveRoom(reservationData)
}

func (service *ReservationService) GetReservationsForUser(user *schemas.UserData) []schemas.ReservationDetails{
	return service.repository.GetReservationsForUser(user)
}
