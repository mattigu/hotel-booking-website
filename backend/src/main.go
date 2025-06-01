package main

import (
	"bd2_projekt/app_err"
	"bd2_projekt/database"
	"bd2_projekt/hotel"
	"context"
	//"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func processQuery(query string){

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), app_err.HTTPStatus(err))
			// Can add behaviour to hide internal app details in error message for unhandled errors
		}
	}
}

func main() {
	// loading environmental variables
	godotenv.Load(".env")

	// connecting to database, which URL is in .env file
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	db, err := database.NewDatabase(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	println("connected to database")
	
	// Test for seed data in db
	tests := getTests();

	fmt.Printf("len=%d cap=%d %v\n", len(tests), cap(tests), tests)

	// setup endpoints
	http.HandleFunc("GET /test/getTest", CORS(getTest))
	http.HandleFunc("GET /hotels", CORS(getHotel))
	http.HandleFunc("POST /test/postTest", CORS(postTest))
		panic(err)
	}

	hotelHandler := hotel.NewHotelHandler(db)

	http.HandleFunc("GET /hotels/getall", CORS(errorHandler(hotelHandler.GetAll)))
	http.HandleFunc("GET /hotels/getbyid", CORS(errorHandler(hotelHandler.GetById)))

	// listen on port 3000
	http.ListenAndServe(":3000", nil)
	db.Close()
}
