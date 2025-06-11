package schemas

type Reservation struct{
	HotelId		int		`json:"hotel_id"`
	RoomId		int		`json:"room_id"`
	StartDate	string	`json:"start_date"`
	EndDate 	string	`json:"end_date"`
	CustomerId	int		`json:"customer_id"`
	PaymentId 	int		`json:"payment_info_id"`
}