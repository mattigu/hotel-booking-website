package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Temporary CORS setup for frontend test
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func getTest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "GET Hello world\n")
}

func postTest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	type testStruct struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	var t testStruct

	err := json.NewDecoder(r.Body).Decode(&t)
	_ = err
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received\n")
	fmt.Print("Message received ", t)
}

func main() {
	godotenv.Load(".env")
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Fprintf(os.Stderr, "connected\n")

	var name string
	err = conn.QueryRow(context.Background(), "select 'HELLO WORLD!'").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(name)

	http.HandleFunc("GET /test/getTest", getTest)
	http.HandleFunc("POST /test/postTest", postTest)

	http.ListenAndServe(":3000", nil)
}
