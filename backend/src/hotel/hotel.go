package hotel

type Hotel struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AddressId    int64  `json:"adress_id"`
	Description  string `json:"description"`
	StarStandard int32  `json:"star_standard"`
}
