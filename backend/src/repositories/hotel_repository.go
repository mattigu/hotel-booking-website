package repositories

import (
	"bd2_projekt/database"
	"bd2_projekt/schemas"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type HotelRepository struct {
	Db *database.Database
}

func (hotelRepository *HotelRepository) GetAll() ([]schemas.Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels"
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query hotels: %w", err)
	}
	defer rows.Close()
	hotels := []schemas.Hotel{}
	for rows.Next() {
		hotel := schemas.Hotel{}
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.AddressId,
			&hotel.Description,
			&hotel.StarStandard)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (hotelRepository *HotelRepository) GetById(id int64) (schemas.Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels where id=@id"
	args := pgx.NamedArgs{
		"id": id,
	}
	hotel := schemas.Hotel{}
	err := hotelRepository.Db.Pool().QueryRow(context.Background(), query, args).Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.AddressId,
		&hotel.Description,
		&hotel.StarStandard)

	if err != nil {
		return schemas.Hotel{}, fmt.Errorf("unable to query hotels: %w", err)
	}
	return hotel, nil
}

func (hotelRepository *HotelRepository) GetHotelsByCity(city string) ([]schemas.HotelOverview, error){
	query := `select h.id, h.name, h.description, h.star_standard 
		from hotels h 
		inner join addresses a on a.id=h.address_id 
		where a.city like @city;`
	
	args := pgx.NamedArgs{
		"city": city,
	}
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var hotels []schemas.HotelOverview
	for rows.Next() {
		var hotel schemas.HotelOverview
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Desc,
			&hotel.Star,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (hotelRepository *HotelRepository) getReservedHotelsId(city string, startDate string, endDate string, ) ([]schemas.HotelInfo, error){
	//query := "SELECT id FROM hotels "
	return nil, nil
}
