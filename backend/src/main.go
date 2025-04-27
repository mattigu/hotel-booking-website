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

type testStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

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

	var t testStruct
	err := json.NewDecoder(r.Body).Decode(&t)
	_ = err
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received\n")
	fmt.Println("Message received ", t)
}

func main() {
	godotenv.Load(".env")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	println("connected to database")

	var name int
	err = conn.QueryRow(context.Background(), "select value from test_table where id=2").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(name)

	// Test for seed data in db
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
	fmt.Printf("len=%d cap=%d %v\n", len(tests), cap(tests), tests)

	http.HandleFunc("GET /test/getTest", getTest)
	http.HandleFunc("POST /test/postTest", postTest)

	http.ListenAndServe(":3000", nil)
}
