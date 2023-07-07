package controller

import (
	"encoding/json"
	"io"
	"net/http"
)


func RegisterControllers(){
	uc:= newUserController()
	http.Handle("/users",*uc)
	http.Handle("/users/",*uc)
}


func encodeResponseAsSJON(data interface{} ,w io.Writer) {
	enc:=json.NewEncoder(w)
	enc.Encode(data)
}