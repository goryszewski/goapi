package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type testwy interface {
	ServeHTTP()
	getAll()
	
}

func RegisterControllers(){
	uc:= newUserController()
	pc:= newPostController()
	http.Handle("/users",*uc)
	http.Handle("/users/",*uc)

	http.Handle("/posts",*pc)
	http.Handle("/posts/",*pc)
}


func encodeResponseAsSJON(data interface{} ,w io.Writer) {
	enc:=json.NewEncoder(w)
	enc.Encode(data)
}