package handlers

import (
	"bd2_projekt/app_err"
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"bd2_projekt/services"
	"encoding/json"
	"strconv"

	"net/http"
)

type HotelHandler struct {
	service *services.HotelService
}

func NewHotelHandler(db *database.Database) *HotelHandler {
	hotelService := services.NewHotelService(db)
	return &HotelHandler{service: hotelService}
}

func (hotelHandler *HotelHandler) GetAll(responseWriter http.ResponseWriter, req *http.Request) error {
	hotels, err := hotelHandler.service.GetAll()
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}

func (hotelHandler *HotelHandler) GetById(responseWriter http.ResponseWriter, req *http.Request) error {
	id := req.PathValue("id")
	guests, err := strconv.Atoi(req.URL.Query().Get("guests"))
	if(err != nil){
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	hotels, err := hotelHandler.service.GetById(int(i), guests)
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}

func (hotelHandler *HotelHandler) GetHotelsSearchQuery(responseWriter http.ResponseWriter, req *http.Request) error {
	var searchQuery schemas.HotelSearchQueryDetails;
	var err error;

	// getting data from the URL
	searchQuery.City = req.URL.Query().Get("city")
	searchQuery.StartDate = req.URL.Query().Get("startdate")
	searchQuery.EndDate = req.URL.Query().Get("enddate")
	searchQuery.Guests, err = strconv.Atoi(req.URL.Query().Get("guests"))

	// geting data for the URL parameters
	hotels, err := hotelHandler.service.GetHotelsSearchQuery(&searchQuery)
	if err != nil { 
		return app_err.WithHTTPStatus(err, http.StatusBadRequest) 
	}
	
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(responseWriter).Encode(hotels)
}
