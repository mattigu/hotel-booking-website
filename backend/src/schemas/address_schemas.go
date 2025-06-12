package schemas

type AddressData struct {
	City 		string `json:"city"`
	Street 		string `json:"street"`
	HouseNumber string `json:"house_number"`
	Country 	string `json:"country"`
}