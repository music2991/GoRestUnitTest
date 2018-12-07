package main

import (
	"GoRestUnitTest/src/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter().StrictSlash(false)
	userWS := services.NewUser()
	router.HandleFunc("/api/user", userWS.GetUser).Methods("GET")
	router.HandleFunc("/api/user", userWS.PostUser).Methods("POST")

	log.Println("Listen and serve")
	http.ListenAndServe(":5000", router)
}

