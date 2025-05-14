package hotel

type Hotel struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AddressId    int    `json:"adress_id"`
	Description  string `json:"description"`
	StarStandard int    `json:"star_standard"`
}
