package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"web/api/model"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r*http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w,r)
		case http.MethodPost:
			uc.post(w,r)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r*http.Request) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Return All users")
	encodeResponseAsSJON(model.GetUsers(),w)
}

func (uc *userController) post(w http.ResponseWriter, r*http.Request) {

	u,err:= uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Problem parse object"))
		return
	}

	u, err = model.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsSJON(u,w)
}

func (uc *userController) parseRequest(r *http.Request)(model.User,error){
	dec:= json.NewDecoder(r.Body)
	var u model.User
	err:= dec.Decode(&u)
	if err != nil {
		return model.User{},err
	}
	return u,nil
}

func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}