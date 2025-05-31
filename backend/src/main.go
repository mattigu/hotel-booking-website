package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
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

}

func main() {
	// loading environmental variables
	godotenv.Load(".env")

	// connecting to database, which URL is in .env file
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
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

	// listen on port 3000
	http.ListenAndServe(":3000", nil)
}
