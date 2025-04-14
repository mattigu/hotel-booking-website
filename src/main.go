package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello w3456orld\n")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}