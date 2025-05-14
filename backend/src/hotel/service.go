package hotel

import (
	"bd2_projekt/database"
)

type hotelService struct {
	r *hotelRepository
}

func newHotelService(db *database.Database) *hotelService {
	r := &hotelRepository{db: db}
	return &hotelService{r: r}
}

func (s *hotelService) getAll() ([]Hotel, error) {
	return s.r.getAll()
}

func (s *hotelService) getById(id int64) (Hotel, error) {
	return s.r.getById(id)
}
