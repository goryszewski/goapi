package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controler struct {
	localchan chan string
}

type Test struct {
	Name string `json:"name"`
}

func (c Controler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p Test

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("err")
	}
	path := r.URL.Path
	rquery := r.URL.RawQuery
	test := r.URL.Query()
	c.localchan <- p.Name
	log.Printf("json: %+v", p)
	log.Printf("url: %+v ", path)
	log.Printf("rquery: %+v ", rquery)
	log.Printf("test: %+v ", test)
	fmt.Fprintf(w, "test api")
}

func newApiControler(testowychan chan string) *Controler {
	return &Controler{localchan: testowychan}
}

func Req01(testowychan chan string) *http.ServeMux {

	apiControler := newApiControler(testowychan)

	req01 := http.NewServeMux()

	req01.Handle("/api", *apiControler)

	return req01
}
