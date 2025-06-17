package schemas

type NewReview struct {
	HotelId 	int 	`json:"hotel_id"`
	Username 	string 	`json:"username"`
	ReviewText 	string 	`json:"review_text"`
	Rating 		int 	`json:"rating"`
}

type ReviewData struct {
	Username 	string 	`json:"username"`
	ReviewText 	string 	`json:"review_text"`
	Rating 		int 	`json:"rating"`
	UploadDate	string	`json:"upload_date"`
}