package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"web/api/api"
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

func buildService(ctx context.Context, service Service, ch chan string) {
	ch <- service.Addr
	server := newService(service)

	fmt.Printf("RUN Service Port: %s in context: [%p] \n", service.Addr, ctx.Value("port"))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error:%p\n", err)
	}
	data := <-ch
	fmt.Printf("LOG1 %s\n", data)
}
func calc(name string) {
	time.Sleep(time.Second * 10)
	fmt.Printf("jakis task w tle %s", name)
}

func main() {
	ctx := context.Background()
	ch := make(chan string, 3)
	ctx = context.WithValue(ctx, "port", ch)
	log.Println("Start Main")
	handler := http.NewServeMux()
	handler.HandleFunc("/", prep(":8088"))
	// load controler
	handler.Handle("/api", api.Req01(ch))

	services := Service{Addr: ":8088", handle: handler}

	go buildService(ctx, services, ch)

	for {
		select {
		case <-ctx.Done():
			log.Println("Context Done")
		case result := <-ch:
			go calc(result)
			log.Printf("Channel : [%s] - \n", result)
		}
		time.Sleep(time.Second * 2)
		fmt.Printf("Loop | chan size: [%x]\n", len(ch))
	}
}
