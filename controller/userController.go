package controller

import (
	// "github.com/yakhyadabo/go-rest-template/src/service"
	"net/http"
	"github.com/yakhyadabo/go-rest-template/src/model"
	"github.com/yakhyadabo/go-rest-template/src/repository"
	"encoding/json"
	u "github.com/yakhyadabo/go-rest-template/src/utils"
	"log"
	// "github.com/gorilla/mux"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	repo := repository.NewUserRepository()
	// service := service.NewUserService(repo)

	repo.Save(user)
	resp := u.Message(true, "success")
	resp["user"] = user
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	login := v.Get("login")

	log.Println("login : ", login)
	repo := repository.NewUserRepository()
	user, error := repo.FindByLogin(string(login))

	if error != nil {
		u.Respond(w, u.Message(false, "There was an error in your repository"))
		return
	}

	if user == nil {
		u.Respond(w, u.Message(false, "No user with login " + login + " found"))
		return
	}

	// data := model.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["user"] = user
	u.Respond(w, resp)
}