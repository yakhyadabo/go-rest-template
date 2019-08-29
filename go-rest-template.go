package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yakhyadabo/go-rest-template/src/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", controller.GetUser).Queries("login", "{login}").Methods("GET")

	router.HandleFunc("/api/employee/new", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee", controller.GetEmployee).Queries("login", "{login}").Methods("GET")

	port := "8000" //localhost

	log.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		log.Print(err)
	}
}
