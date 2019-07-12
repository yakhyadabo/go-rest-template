package main

import (
	"github.com/gorilla/mux"
	"github.com/yakhyadabo/go-rest-template/src/controller"
	"os"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", controller.GetUser).Queries("login", "{login}",).Methods("GET")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" 
	}

	log.Println(port)

	err := http.ListenAndServe(":" + port, router) 
	if err != nil {
		log.Print(err)
	}
}
