package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controler struct {
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
	log.Printf("json: %+v", p)
	fmt.Fprintf(w, "test api")
}

func newApiControler() *Controler {
	return &Controler{}
}

func Req01() *http.ServeMux {

	apiControler := newApiControler()

	req01 := http.NewServeMux()

	req01.Handle("/api", *apiControler)

	return req01
}
