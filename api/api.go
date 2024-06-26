package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	model "web/api/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controler struct {
	ctx   context.Context
	db    *model.REDIS
	mongo *model.DB
}

type Test struct {
	Name string `json:"name"`
}

func (c Controler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p Test
	var val string = ""
	var err error

	var code int

	path := r.URL.Path
	rquery := r.URL.RawQuery
	test := r.URL.Query()
	log.Printf("url: %+v ", path)
	log.Printf("rquery: %+v ", rquery)
	log.Printf("test: %+v ", test)
	log.Printf("json: %+v", p)
	collection := c.mongo.Get()
	if r.Method == "POST" {

		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println("err")
		}
		val = p.Name
		err = c.db.Set(c.ctx, "key", val, 0).Err()
		if err != nil {
			panic(err)
		}
		code = http.StatusCreated

		_, err := collection.InsertOne(c.ctx, &p)
		log.Printf(err.Error())

	} else if r.Method == "GET" {
		val, err = c.db.Get(c.ctx, "key").Result()
		code = http.StatusAccepted

	} else if r.Method == "DELETE" {
		c.db.Del(c.ctx, "key")
		code = http.StatusAccepted
		filter := bson.D{primitive.E{}}
		collection.DeleteMany(c.ctx, filter)

	} else {
		log.Printf("Bad request")
		code = http.StatusBadGateway
	}

	log.Printf("Method: %+v", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(val)

}

func newApiControler(ctx context.Context) *Controler {
	rdb := model.NewREdis(ctx)
	client := model.NewDB(ctx)
	return &Controler{ctx: ctx, db: rdb, mongo: client}
}

func Req01(ctx context.Context) *http.ServeMux {

	apiControler := newApiControler(ctx)

	req01 := http.NewServeMux()

	req01.Handle("/api", *apiControler)

	return req01
}
