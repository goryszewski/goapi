package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var PORT string = ":8083"

func prep(port string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Page Test API PORT: [%s]\n", port)
		fmt.Fprintf(w, "Page Test API PORT: [%s]\n", port)
	}
}

func newService(addr string, handler http.Handler) *http.Server {
	return &http.Server{Addr: addr, Handler: handler}
}

func buildService(port string, ch chan string, handler http.Handler) {
	ch <- port
	server := newService(port, handler)

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

	services := []string{":8084", ":8083", ":8085"}

	for _, item := range services {
		handler := http.NewServeMux()
		handler.HandleFunc("/", prep(item))
		go buildService(item, ch, handler)
	}

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("Loop | chan size: [%x]\n", len(ch))
	}
}
