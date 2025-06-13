package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
//	"fmt"
//	"os"
	"github.com/jackc/pgx/v5"
//	"strconv"
)

type ReservationRepository struct{
	Db *database.Database
}

func (reservationRepository *ReservationRepository) ReserveRoom(reservationData *schemas.Reservation) error{
	query := `INSERT INTO reservations ("id", "customer_id", "hotel_id", "room_id", "start_date", "end_date", "payment_info_id") 
		VALUES ((SELECT max(id) + 1 FROM reservations), @customerId, @hotelId, @roomId, @startDate, @endDate, @paymentInfoId)`

	args := pgx.NamedArgs{
		"customerId": reservationData.CustomerId,
		"hotelId": reservationData.HotelId,
		"roomId": reservationData.RoomId,
		"startDate": reservationData.StartDate,
		"endDate": reservationData.EndDate,
		"paymentInfoId": reservationData.PaymentId,
	}

	reservationRepository.Db.Pool().Query(context.Background(), query, args)

	return nil
}
