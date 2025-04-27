package main

import (
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

// Temporary CORS setup for frontend test
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func main() {
	godotenv.Load(".env")
	db, err := database.NewDatabase(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	hotelHandler := hotel.NewHotelHandler(db)

	http.HandleFunc("GET /hotels/getall", CORS(hotelHandler.GetAll))

	http.ListenAndServe(":3000", nil)
	db.Close()
}
