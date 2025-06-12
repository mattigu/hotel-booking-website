package schemas

type Hotel struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AddressId    int64  `json:"adress_id"`
	Description  string `json:"description"`
	StarStandard int32  `json:"star_standard"`
}

type TestStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type HotelOverview struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
	Star int64 `json:"star_standard"`
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
