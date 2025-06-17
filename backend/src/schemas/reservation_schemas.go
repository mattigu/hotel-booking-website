package schemas

type Reservation struct{
	HotelId		int		`json:"hotel_id"`
	RoomId		int		`json:"room_id"`
	StartDate	string	`json:"start_date"`
	EndDate 	string	`json:"end_date"`
	Customer 	UserData `json:"customer"`
	PaymentId 	PaymentInfo `json:"payment_info"`
	Addons	 	[]int 	`json:"addons"`
}

type ReservationDetails struct{
	HotelId		int		`json:"hotel_id"`
	RoomId		int		`json:"room_id"`
	StartDate	string	`json:"start_date"`
	EndDate 	string	`json:"end_date"`
}