package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yakhyadabo/go-rest-template/model"
	"github.com/yakhyadabo/go-rest-template/repository"
	u "github.com/yakhyadabo/go-rest-template/utils"
)

var CreateEmployee = func(w http.ResponseWriter, r *http.Request) {

	employee := &model.Employee{}
	err := json.NewDecoder(r.Body).Decode(employee)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	repo := repository.NewEmployeeRepository()
	// service := service.NewEmployeeService(repo)

	repo.Save(employee)
	resp := u.Message(true, "success")
	resp["employee"] = employee
	u.Respond(w, resp)
}

var GetEmployee = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	login := v.Get("login")

	log.Println("login : ", login)
	repo := repository.NewEmployeeRepository()
	employee, error := repo.FindByLogin(string(login))

	if error != nil {
		u.Respond(w, u.Message(false, "There was an error in your repository"))
		return
	}

	if employee == nil {
		u.Respond(w, u.Message(false, "No employee with login "+login+" found"))
		return
	}

	// data := model.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["employee"] = employee
	u.Respond(w, resp)
}
