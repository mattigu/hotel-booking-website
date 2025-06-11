package main

import (
//	"bd2_projekt/app_err"
//	"bd2_projekt/database"
	"bd2_projekt/hotel"
//	"context"
//	"encoding/json"
//	"fmt"
	"net/http"
//	"os"

	"github.com/joho/godotenv"
//	"github.com/jackc/pgx/v5"
)

func main() {
	// loading environmental variables
	godotenv.Load(".env")

	// connecting to the database
	db := connectToDatabase();
	println("connected to database")
	
	// creating the hotel handler
	hotelHandler := hotel.NewHotelHandler(db)

	// create all endpoints
	http.HandleFunc("GET /hotels", CORS(errorHandler(hotelHandler.GetHotelsByCity)))
	http.HandleFunc("GET /hotels/getall", CORS(errorHandler(hotelHandler.GetAll)))
	http.HandleFunc("GET /hotels/getbyid", CORS(errorHandler(hotelHandler.GetById)))

	// listen on port 3000
	http.ListenAndServe(":3000", nil)
	db.Close()
}
