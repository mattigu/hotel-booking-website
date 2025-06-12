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

func (hotelRepository *HotelRepository) GetHotelsSearchQuery(searchQuery *schemas.HotelSearchQueryDetails) ([]schemas.HotelInfo, error){
	query := `SELECT h.id, h.name, h.star_standard, r.single_bed_num, r.double_bed_num
		FROM hotels h 
			INNER JOIN addresses a on a.id=h.address_id
			INNER JOIN rooms r on r.hotel_id=h.id
		WHERE a.city like @city AND r.single_bed_num + r.double_bed_num * 2 = @guests;`
	
	args := pgx.NamedArgs{
		"city": searchQuery.City,
		"guests": searchQuery.Guests,
	}
	rows, err := hotelRepository.Db.Pool().Query(context.Background(), query, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var hotels []schemas.HotelInfo
	for rows.Next() {
		var hotel schemas.HotelInfo
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Star,
			&hotel.SingleBeds,
			&hotel.DoubleBeds,
		)
		hotel.Price = "100.99";
		hotel.PhotoUrl = "https://content.r9cdn.net/rimg/kimg/4a/83/41fc6b329baa5c28.jpg?width=1200&height=630&crop=true";
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
