package controller

import (
	//"github.com/yakhyadabo/go-rest-template/src/service"
	"net/http"
	"github.com/yakhyadabo/go-rest-template/src/model"
	//"github.com/yakhyadabo/go-rest-template/src/repository"
	"encoding/json"
	u "github.com/yakhyadabo/go-rest-template/src/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	//userId := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	// user := model.NewUser("userId","email")
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

 // repo := repository.NewUserRepository()
  // service := service.NewUserService(repo)

	//user.id = userId
	// resp :=  repo.Save(user)
	resp := u.Message(true, "success")
	resp["user"] = user
	u.Respond(w, resp)
}

/*
var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := model.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}*/