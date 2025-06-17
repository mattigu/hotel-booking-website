package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
	"strconv"
	"time"
	"math"
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

func (repository *ReservationRepository) getRoomPrice(roomId int, days int) int{
	query := `SELECT base_price
	FROM rooms
	WHERE id=@id`

	args := pgx.NamedArgs{
		"id": roomId,
	}
	rows, err := repository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var cost int;
	if(rows.Next()){
		err := rows.Scan(&cost)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		return cost * days
	}

	return 0
}

func (repository *ReservationRepository) calculatePrice(addons []int) int{
	var addons_string string;

	for i, value := range addons {
		if i != 0 {
			addons_string = addons_string + ", "
		}
		addons_string = addons_string + strconv.Itoa(value)
	}

	query := `SELECT sum(price)
	FROM reservation_addons
	WHERE id in (` + addons_string + `)`

	rows, err := repository.Db.Pool().Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var sum int;
	if(rows.Next()){
		err := rows.Scan(&sum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		return sum
	}

	return 0
}

func (repository *ReservationRepository) addUser(user *schemas.UserData) error{
	query := `INSERT INTO customers ("name", "surname", "phone_number") 
		VALUES (@name, @surname, @phoneNumber)`

	args := pgx.NamedArgs{
		"name": user.Name,
		"surname": user.Surname,
		"phoneNumber": user.PhoneNumber,
	}

	repository.Db.Pool().Exec(context.Background(), query, args)

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
	now := time.Now()
	currentDate := now.Truncate(24 * time.Hour)
	formattedDate := currentDate.Format("2006-01-02")

	query := `INSERT INTO payments ("id", "payment_type", "payment_data", "due_date", "amount", "fulfilled")
	VALUES((SELECT max(id) + 1 FROM payments), '`+paymentInfo.PaymentType+`', '`+
	paymentInfo.PaymentData+`', '`+ formattedDate +`', '` + 
	strconv.Itoa(paymentInfo.Amount)+`', true)
	RETURNING id;`

	var paymentId int;
	err := repository.Db.Pool().QueryRow(context.Background(), query).Scan(&paymentId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}

	return paymentId;
}

func (repository *ReservationRepository) ReserveRoom(reservationData *schemas.Reservation) error{
	userId := repository.addUserIfAbsent(&reservationData.Customer)

	timeStart, _ := time.Parse("2006-01-02", reservationData.StartDate)
	timeEnd, _ := time.Parse("2006-01-02", reservationData.EndDate)

	duration := timeEnd.Sub(timeStart);

	reservationData.PaymentId.Amount = repository.calculatePrice(reservationData.Addons) +
		repository.getRoomPrice(reservationData.RoomId, int(math.Round(duration.Hours()/24)))
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

	repository.Db.Pool().Exec(context.Background(), query, args)

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
