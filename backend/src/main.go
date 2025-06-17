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
	// "github.com/jackc/pgx/v5"
)

func main() {
	// loading environmental variables
	godotenv.Load(".env")

	// connecting to the database
	db := connectToDatabase()
	println("connected to database")

	// creating all the handlers
	hotelHandler := handlers.NewHotelHandler(db)
	reservationHandler := handlers.NewReservationHandler(db)
	reviewHandler := handlers.NewReviewHandler(db)

	http.HandleFunc("OPTIONS /", CORS(func(w http.ResponseWriter, r *http.Request) {}))

	// create all get endpoints
	http.HandleFunc("GET /get/hotels", CORS(errorHandler(hotelHandler.GetHotelsSearchQuery)))
	http.HandleFunc("GET /get/hotel/{id}", CORS(errorHandler(hotelHandler.GetById)))
	http.HandleFunc("GET /hotels/getall", CORS(errorHandler(hotelHandler.GetAll)))
	http.HandleFunc("GET /get/reservations", CORS(errorHandler(reservationHandler.GetReservationsFor)))
	http.HandleFunc("GET /get/configuration", CORS(errorHandler(hotelHandler.GetRoomConfigurations)))

	// create all post endpoints
	http.HandleFunc("POST /reserve/room", CORS(errorHandler(reservationHandler.ReserveRoom)))
	http.HandleFunc("POST /new/review", CORS(errorHandler(reviewHandler.PostReview)))

	// listen on port 3000
	http.ListenAndServe(":3000", nil)

	// close the database connection
	db.Close()
}
