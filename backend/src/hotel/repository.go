package hotel

import (
	"bd2_projekt/database"
	"context"
	"fmt"
	"os"
)

type hotelRepository struct {
	db *database.Database
}

func (r *hotelRepository) getAll() ([]Hotel, error) {
	rows, err := r.db.Pool().Query(context.Background(), "select id, name from hotels")
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
			&hotel.AdressId,
			&hotel.Description,
			&hotel.StarStandard)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}
