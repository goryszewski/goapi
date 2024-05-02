package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var PORT string = ":8083"

type Service struct {
	Addr   string
	handle *http.ServeMux
}

func prep(port string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Page Test API PORT: [%s]\n", port)
		fmt.Fprintf(w, "Page Test API PORT: [%s]\n", port)
	}
}

func newService(service Service) *http.Server {
	return &http.Server{Addr: service.Addr, Handler: service.handle}
}

func buildService(service Service, ch chan string) {
	ch <- service.Addr
	server := newService(service)

	fmt.Printf("RUN Service Port: %s \n", service.Addr)
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
	handler := http.NewServeMux()
	handler.HandleFunc("/", prep(":8088"))
	// load controler
	handler.Handle("/api", req01())

	services := Service{Addr: ":8088", handle: handler}

	go buildService(services, ch)

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("Loop | chan size: [%x]\n", len(ch))
	}
}
