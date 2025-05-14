package hotel

import (
	"bd2_projekt/app_err"
	"bd2_projekt/database"
	"encoding/json"
	"strconv"

	"net/http"
)

type HotelHandler struct {
	s *hotelService
}

func NewHotelHandler(db *database.Database) *HotelHandler {
	s := newHotelService(db)
	return &HotelHandler{s: s}
}

func (h *HotelHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	hotels, err := h.s.getAll()
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(w).Encode(hotels)
}

func (h *HotelHandler) GetById(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	hotels, err := h.s.getById(int64(i))
	if err != nil {
		return app_err.WithHTTPStatus(err, http.StatusBadRequest)
	}
	// w.WriteHeader(http.Status...) Change status code here if needed
	return json.NewEncoder(w).Encode(hotels)
}
