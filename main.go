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

func newService(addr string) *http.Server {
	return &http.Server{Addr: addr}
}

func main() {

	go func(port string) {
		server := newService(port)
		fmt.Printf("RUN Service Port: %s \n", port)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error:%p\n", err)
		}
		fmt.Printf("LOG1\n")

	}(":8084")

	log.Println("Start Main")
	Local_Router()
	initweb(PORT)

}
