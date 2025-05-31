package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"net/http"
	//"os"

	//"github.com/jackc/pgx/v5"
	//"github.com/joho/godotenv"
)

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET Hello world\n")
}

func postTest(w http.ResponseWriter, r *http.Request) {
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

func getHotel(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context();
	city := r.URL.Query().Get("city")
	hotels := getHotelsForCity(city)
	jsonData, _ := json.Marshal(hotels)
	fmt.Fprintf(w, string(jsonData))
}