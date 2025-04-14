package main

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"github.com/jackc/pgx/v5"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world\n")
}

func main(){
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgresql://root:root@db:5432/bd2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Fprintf(os.Stderr, "connected\n")

	var name string
	err = conn.QueryRow(context.Background(), "select 'HELLO WORLD!'").Scan(&name)
	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(name)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}