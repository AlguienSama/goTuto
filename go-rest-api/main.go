package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type contact struct {
	ID     int    `json:ID`
	Name   string `json:Name`
	Number string `json:PhoneNumber`
}

type allContacts []contact

var contacts = allContacts{
	{
		ID:     1,
		Name:   "Persona 1",
		Number: "644093423",
	},
	{
		ID:     2,
		Name:   "Persona 2",
		Number: "644453656",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main Route")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
}
