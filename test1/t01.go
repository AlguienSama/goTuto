package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type nums struct {
	num1 int
	num2 int
}

func homePage(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	var t nums
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	//fmt.Println(r)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
