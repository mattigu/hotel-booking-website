package schemas

type Hotel struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AddressId    int64  `json:"adress_id"`
	Description  string `json:"description"`
	StarStandard int32  `json:"star_standard"`
}

type HotelInfo struct {
	Id   		int64  	`json:"id"`
	Name 		string 	`json:"name"`
	Price 		string 	`json:"price"`
	Star 		int64 	`json:"star_standard"`
	SingleBeds 	int32 	`json:"sigle_beds"`
	DoubleBeds 	int32 	`json:"double_beds"`
	PhotoUrl	string	`json:"photo_url"`
}

type HotelSearchQueryDetails struct {
	City		string
	StartDate	string
	EndDate		string
	Guests		int
}

type HotelSpecificData struct {
	Id   			int  		`json:"id"`
	Name 			string 		`json:"name"`
	PhotoUrl		string		`json:"photo_url"`
	Description  	string 		`json:"description"`
	Address 		AddressData	`json:"address"`
	Amenities 		[]Amenities	`json:"amenities"`
	Reviews 		[]ReviewData `json:"reviews"`
}
