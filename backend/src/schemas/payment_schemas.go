package schemas

type PaymentInfo struct {
	PaymentType string `json:"payment_type"`
	PaymentData string `json:"payment_data"`
	Amount		int `json:"amount"`
}