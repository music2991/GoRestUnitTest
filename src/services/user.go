package services

import (
	"GoRestUnitTest/src/model"
	"encoding/json"
	"net/http"
)

type userWS struct {
}

func NewUser() userWS {
	if user == nil {
		user = make(map[string]model.User)
	}
	return userWS{}
}

var user map[string]model.User

func (userWS) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userListResult, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userListResult)
}

func (userWS) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	newUser := model.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user[newUser.Name] = newUser

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
