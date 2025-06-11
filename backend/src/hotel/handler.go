package hotel

import (
	"bd2_projekt/app_err"
	"bd2_projekt/database"
	"encoding/json"
	"strconv"

	"net/http"
)

type HotelHandler struct {
	service *HotelService
}

func NewHotelHandler(db *database.Database) *HotelHandler {
	hotelService := newHotelService(db)
	return &HotelHandler{service: hotelService}
}

func (hotelHandler *HotelHandler) GetAll(responseWriter http.ResponseWriter, req *http.Request) error {
	hotels, err := hotelHandler.service.getAll()
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}

func (hotelHandler *HotelHandler) GetById(responseWriter http.ResponseWriter, req *http.Request) error {
	id := req.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	hotels, err := hotelHandler.service.getById(int64(i))
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}

func (hotelHandler *HotelHandler) GetHotelsByCity(responseWriter http.ResponseWriter, req *http.Request) error {
	hotels, err := hotelHandler.service.getHotelsByCity(req.URL.Query().Get("city"))
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}
