package main

import (
	"fmt"
	"log"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page Test API")
}

func main() {

	http.HandleFunc("/", page)
	log.Fatal(http.ListenAndServe(":8083", nil))
	fmt.Println("test")
}
