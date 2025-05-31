package main

import (
	"context"
	//"encoding/json"
	"fmt"
	//"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	//"github.com/joho/godotenv"
)

var conn *pgx.Conn;
var err error;

func getTests()([]testStruct){
	rows, err := conn.Query(context.Background(), "select id, name from teststruct")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var tests []testStruct
	for rows.Next() {
		var test testStruct
		err := rows.Scan(
			&test.Id,
			&test.Name,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning rows %v\n", err)
		}
		tests = append(tests, test)
	}
	return tests
}

func getHotelsForCity(city string) ([]hotelOverview){
	query := `select h.id, h.name, h.description, h.star_standard 
		from hotels h 
		inner join addresses a on a.id=h.address_id 
		where a.city like '`+city+`';`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve rows from db %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var hotels []hotelOverview
	for rows.Next() {
		var hotel hotelOverview
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
	return hotels
}

