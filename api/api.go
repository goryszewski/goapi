package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Controler struct {
	ctx context.Context
	db  *redis.Client
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

	err = c.db.Set(c.ctx, "key", p.Name, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := c.db.Get(c.ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	log.Printf("json: %+v", p)
	log.Printf("url: %+v ", path)
	log.Printf("rquery: %+v ", rquery)
	log.Printf("test: %+v ", test)
	fmt.Fprintf(w, val)
}

func newApiControler(ctx context.Context) *Controler {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &Controler{ctx: ctx, db: rdb}
}

func Req01(ctx context.Context) *http.ServeMux {

	apiControler := newApiControler(ctx)

	req01 := http.NewServeMux()

	req01.Handle("/api", *apiControler)

	return req01
}
