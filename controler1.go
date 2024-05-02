package main

import (
	"fmt"
	"log"
	"net/http"
)

type Controler struct {
}

func (c Controler) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	log.Printf("test\n")
	fmt.Fprintf(r, "test api")
}

func newApiControler() *Controler {
	return &Controler{}
}

func req01() *http.ServeMux {

	apiControler := newApiControler()

	req01 := http.NewServeMux()

	req01.Handle("/api", *apiControler)

	return req01
}
