package hotel

import (
	"bd2_projekt/database"
)

type HotelService struct {
	repository *HotelRepository
}

func newHotelService(db *database.Database) *HotelService {
	hotelRepository := &HotelRepository{db: db}
	return &HotelService{repository: hotelRepository}
}

func (hotelService *HotelService) getAll() ([]Hotel, error) {
	return hotelService.repository.getAll()
}

func (hotelService *HotelService) getById(id int64) (Hotel, error) {
	return hotelService.repository.getById(id)
}

func (hotelService *HotelService) getHotelsByCity(city string) ([]HotelOverview, error) {
	return hotelService.repository.getHotelsByCity(city)
}
