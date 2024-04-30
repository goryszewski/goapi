package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var PORT string = ":8083"

func page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page Test API")
}

func Local_Router() {
	log.Println("Load Router")
	http.HandleFunc("/", page)
}

func newService(addr string) *http.Server {
	return &http.Server{Addr: addr}
}

func buildService(port string, ch chan string) {
	ch <- port
	server := newService(port)
	fmt.Printf("RUN Service Port: %s \n", port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error:%p\n", err)
	}
	data := <-ch
	fmt.Printf("LOG1 %s\n", data)
}

func main() {
	ch := make(chan string, 3)
	log.Println("Start Main")
	Local_Router()
	services := []string{":8084", ":8083", ":8085"}
	for _, item := range services {

		go buildService(item, ch)
	}

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("Loop | chan size: [%x]\n", len(ch))
	}
}
