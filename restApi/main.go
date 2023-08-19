package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type Profile struct{
	Department string `json:"department"`
	Designation string `json:"designation"`
	Employee User string `json: "employee"`
}

func main_abc() {
	router:=mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")

	http.ListenandServer(":5000", router)
}