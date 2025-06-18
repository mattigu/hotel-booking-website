package services

import (
	"bd2_projekt/database"
	"bd2_projekt/repositories"
	"bd2_projekt/schemas"
)

type HotelService struct {
	repository *repositories.HotelRepository
}

func NewHotelService(db *database.Database) *HotelService {
	hotelRepository := &repositories.HotelRepository{Db: db}
	return &HotelService{repository: hotelRepository}
}

func (hotelService *HotelService) GetAll() ([]schemas.Hotel, error) {
	return hotelService.repository.GetAll()
}

func (hotelService *HotelService) GetById(id int, guests int) (schemas.HotelSpecificData, error) {
	return hotelService.repository.GetById(id, guests)
}

func (hotelService *HotelService) GetHotelsSearchQuery(searchQuery *schemas.HotelSearchQueryDetails) ([]schemas.HotelInfo, error) {
	return hotelService.repository.GetHotelsSearchQuery(searchQuery)
}

func (service *HotelService) GetRoomConfigurations(hotelId int, guests int, startDate string, endDate string) []schemas.RoomConfiguration{
	val, _ := service.repository.GetRoomConfigurations(hotelId, guests, startDate, endDate)
	return val
}

func (service *HotelService) GetAddons(hotelId int) []schemas.AddonData{
	val, _ := service.repository.GetAddons(hotelId)
	return val
}