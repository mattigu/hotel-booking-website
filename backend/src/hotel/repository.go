package hotel

import (
	"bd2_projekt/database"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type hotelRepository struct {
	db *database.Database
}

func (r *hotelRepository) getAll() ([]Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels"
	rows, err := r.db.Pool().Query(context.Background(), query)
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

func (r *hotelRepository) getById(id int64) (Hotel, error) {
	query := "select id, name, address_id, description, star_standard from hotels where id=@id"
	args := pgx.NamedArgs{
		"id": id,
	}
	hotel := Hotel{}
	err := r.db.Pool().QueryRow(context.Background(), query, args).Scan(
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
