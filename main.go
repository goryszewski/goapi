package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"web/api/controller"
)
func agrs() int {
	var nFlag = flag.Int("port",8080,"Port binding service")

	flag.Parse()
	return *nFlag
}

func main() {

	args := os.Args
	for _ , v := range args {
		fmt.Println(v)
	}


	port := agrs()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

		logger.Println("RUN api server")
	controller.RegisterControllers()
		logger.Println("Register Controller")
	settings:= fmt.Sprintf(":%v",port)
		logger.Printf("Start api server %v\n",settings)
	http.ListenAndServe(settings,nil)
}