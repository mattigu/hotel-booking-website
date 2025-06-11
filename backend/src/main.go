package main

import (
//	"bd2_projekt/app_err"
//	"bd2_projekt/database"
	"bd2_projekt/handlers"
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
	hotelHandler := handlers.NewHotelHandler(db)

	// create all endpoints
	http.HandleFunc("GET /get/hotels", CORS(errorHandler(hotelHandler.GetHotelsByCity)))
	http.HandleFunc("GET /get/hotel/{id}", CORS(errorHandler(hotelHandler.GetById)))
	http.HandleFunc("GET /hotels/getall", CORS(errorHandler(hotelHandler.GetAll)))

	// listen on port 3000
	http.ListenAndServe(":3000", nil)

	// close the database connection
	db.Close()
}
