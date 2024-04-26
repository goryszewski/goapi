package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT string = ":8083"

func page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page Test API")
}

func Local_Router() {
	log.Println("Load Router")
	http.HandleFunc("/", page)
}

func initweb(port string) {
	log.Println("Start Main")
	err := http.ListenAndServe(port, nil)

	log.Fatal(err)
}

func main() {

	log.Println("Start Main")
	Local_Router()
	initweb(PORT)

}
