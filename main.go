package main

import (
	"log"
	"net/http"
	"os"
	"web/api/controller"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Start api server")
	controller.RegisterControllers()
	logger.Println("Register Controller")
	http.ListenAndServe(":8080",nil)
}