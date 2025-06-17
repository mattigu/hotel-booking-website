package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
//	"strconv"
)

type ReservationRepository struct{
	Db *database.Database
}

// returns the user id or -1 if user does not exist
func (repository *ReservationRepository) getUserId(user *schemas.UserData) (int, error){
	query := `SELECT id
	FROM customers
	WHERE name like @name and surname like @surname and phone_number like @phoneNumber;`

	args := pgx.NamedArgs{
		"name": user.Name,
		"surname": user.Surname,
		"phoneNumber": user.PhoneNumber,
	}
	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var id int;
	if(rows.Next()){
		err := rows.Scan(&id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		return id, nil
	}

	return -1, nil
}

func (repository *ReservationRepository) addUser(user *schemas.UserData) error{
	query := `INSERT INTO customers ("name", "surname", "phone_number") 
		VALUES (@name, @surname, @phoneNumber)`

	args := pgx.NamedArgs{
		"name": user.Name,
		"surname": user.Surname,
		"phoneNumber": user.PhoneNumber,
	}

	repository.Db.Pool().Query(context.Background(), query, args)

	return nil
}

func (repository *ReservationRepository) addUserIfAbsent(user *schemas.UserData) int{
	userId, _ := repository.getUserId(user)
	if(userId == -1){
		repository.addUser(user);
		userId, _ = repository.getUserId(user)
	}
	return userId;
}

func (repository *ReservationRepository) addPaymentInfo (paymentInfo *schemas.PaymentInfo) int {
	query := `INSERT INTO payments ("id", "payment_type", "payment_data", "due_date", "amount", "fulfilled")
	VALUES((SELECT max(id) + 1 FROM payments), '`+paymentInfo.PaymentType+`', '`+paymentInfo.PaymentData+`', '1970-01-01', '` + paymentInfo.Amount+`', true)
	RETURNING id;`

	rows, err := repository.Db.Pool().Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var paymentId int;
	for rows.Next(){
		err := rows.Scan(
			&paymentId,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
	}
	return paymentId;
}

func (repository *ReservationRepository) ReserveRoom(reservationData *schemas.Reservation) error{
	userId := repository.addUserIfAbsent(&reservationData.Customer)
	paymentId := repository.addPaymentInfo(&reservationData.PaymentId)

	query := `INSERT INTO reservations ("id", "customer_id", "hotel_id", "room_id", "start_date", "end_date", "payment_info_id") 
		VALUES ((SELECT max(id) + 1 FROM reservations), @customerId, @hotelId, @roomId, @startDate, @endDate, @paymentInfoId)`

	args := pgx.NamedArgs{
		"customerId": userId,
		"hotelId": reservationData.HotelId,
		"roomId": reservationData.RoomId,
		"startDate": reservationData.StartDate,
		"endDate": reservationData.EndDate,
		"paymentInfoId": paymentId,
	}

	repository.Db.Pool().Query(context.Background(), query, args)

	return nil
}

func (repository *ReservationRepository) GetReservationsForUser(user *schemas.UserData) []schemas.ReservationDetails{
	user_id, _ := repository.getUserId(user)
	query := `SELECT hotel_id, room_id, start_date::text, end_date::text
	FROM reservations
	WHERE customer_id=@id`

	args := pgx.NamedArgs{
		"id": user_id,
	}

	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var reservations []schemas.ReservationDetails
	for rows.Next() {
		var reservation schemas.ReservationDetails
		err := rows.Scan(
			&reservation.HotelId,
			&reservation.RoomId,
			&reservation.StartDate,
			&reservation.EndDate,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		reservations = append(reservations, reservation)
	}
	return reservations
}
