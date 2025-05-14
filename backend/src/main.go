package main

import (
	"bd2_projekt/app_err"
	"bd2_projekt/database"
	"bd2_projekt/hotel"
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	println("here")
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
	godotenv.Load(".env")
	db, err := database.NewDatabase(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	hotelHandler := hotel.NewHotelHandler(db)

	http.HandleFunc("GET /hotels/getall", CORS(errorHandler(hotelHandler.GetAll)))
	http.HandleFunc("GET /hotels/getbyid/{id}", CORS(errorHandler(hotelHandler.GetById)))

	http.ListenAndServe(":3000", nil)
	db.Close()
}
