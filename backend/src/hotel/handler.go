package hotel

import (
	"bd2_projekt/database"
	"encoding/json"
	"net/http"
)

type HotelHandler struct {
	s *hotelService
}

func NewHotelHandler(db *database.Database) *HotelHandler {
	s := newHotelService(db)
	return &HotelHandler{s: s}
}

func (h *HotelHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	hotels, err := h.s.getAll()

	// Think about this later !!
	// https://www.reddit.com/r/golang/comments/1128rt5/error_handling_in_http_handlefuncs/

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(hotels)
}
