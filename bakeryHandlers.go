package main

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
)

func OrderBreads(writer http.ResponseWriter, request *http.Request) {
	log.Println("Order new Breads")
	initHeaders(writer)

	var requestedBreads Breads
	_ = json.NewDecoder(request.Body).Decode(&requestedBreads)

	//Try to find in stock the breads.
	Breads, idx := FindBreadFromBakery(requestedBreads.Pizza)

	//If idx is inferior than 0, we need to produce new breads
	if idx < 0 {
		//Initialize a client
		client := &http.Client{}
		//Prepare a PUT Request to http://localhost:8000/bakery
		request, _ := http.NewRequest(http.MethodPut, "http://localhost:8000/bakery", nil)
		//Send the request
		client.Do(request)
		//"Reload" the breads
		breads, idx = FindBreadFromBakery(requestedBreads.Pizza)
	}

	requestedBreads = breads
	//Removes the barrel from the stock
	bakeryBreads = append(bakeryBreads[:idx], bakeryBreads[idx+1:]...)

	json.NewEncoder(writer).Encode(requestedBreads)
}

func ProduceBreads(writer http.ResponseWriter, request *http.Request) {
	log.Println("Producing new Breads")

	bakeryBreads = append(bakeryBreads,
		Breads{&pizza[0], 1000, time.Now()},
		Breads{&pizza[1], 5000, time.Now()},
		Breads{&pizza[2], 3000, time.Now()})
}
