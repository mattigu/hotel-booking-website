package hotel

import (
	"bd2_projekt/database"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type HotelRepository struct {
	db *database.Database
}

func (hotelRepository *HotelRepository) getAll() ([]Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels"
	rows, err := hotelRepository.db.Pool().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query hotels: %w", err)
	}
	defer rows.Close()
	hotels := []Hotel{}
	for rows.Next() {
		hotel := Hotel{}
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

func (hotelRepository *HotelRepository) getById(id int64) (Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels where id=@id"
	args := pgx.NamedArgs{
		"id": id,
	}
	hotel := Hotel{}
	err := hotelRepository.db.Pool().QueryRow(context.Background(), query, args).Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.AddressId,
		&hotel.Description,
		&hotel.StarStandard)

	if err != nil {
		return Hotel{}, fmt.Errorf("unable to query hotels: %w", err)
	}
	return hotel, nil
}

func (hotelRepository *HotelRepository) getHotelsByCity(city string) ([]HotelOverview, error){
	query := `select h.id, h.name, h.description, h.star_standard 
		from hotels h 
		inner join addresses a on a.id=h.address_id 
		where a.city like @city;`
	
	args := pgx.NamedArgs{
		"city": city,
	}
	rows, err := hotelRepository.db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var hotels []HotelOverview
	for rows.Next() {
		var hotel HotelOverview
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
