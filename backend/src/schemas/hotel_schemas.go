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

type RoomConfiguration struct {
	SingleBeds 	int	`json:"single_beds"`
	DoubleBeds 	int	`json:"double_beds"`
	Price		int	`json:"price"`
}

type Addons struct {
	Name string `json:"name"`
	Price string `json:"price"`
}

type HotelSpecificData struct {
	Id   			int  		`json:"id"`
	Name 			string 		`json:"name"`
	PhotoUrl		string		`json:"photo_url"`
	StarStandard 	int  		`json:"star_standard"`
	AvgRating	 	float32		`json:"avg_rating"`
	RatingsCount 	int			`json:"rating_count"`
	Description  	string 		`json:"description"`
	Address 		AddressData	`json:"address"`
	Amenities 		[]Amenities	`json:"amenities"`
	Reviews 		[]ReviewData `json:"reviews"`
	RoomConfigurations []RoomConfiguration `json:"room_configurations"`
	Addons			[]Addons	`json:"addons"`
}

type AddonData struct {
	Id		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Price 	int 	`json:"price"`
}
