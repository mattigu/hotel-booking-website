package main

import (
	"context"
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

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	println("info from handler")

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

	http.HandleFunc("/", handler)
	http.HandleFunc("/get", handler)
	http.ListenAndServe(":3000", nil)
}
