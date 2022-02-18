package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"bytes"
)

const plateQuantity = 50

func GetInfo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get infos about bpizzas")
	initHeaders(writer)
	json.NewEncoder(writer).Encode(pizza)
}

func GetPizzaInfo(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	//Converts the id parameter from a string to an int
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err == nil {
		log.Println("Get info about pizza id #", id)

		//Retrieves the infos about the pizza
		pizza := FindPizzaByID(id)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		log.Fatal(err.Error())
	}
}

func OrderPizza(writer http.ResponseWriter, request *http.Request) {
	log.Println("Order a pizza")
	initHeaders(writer)
	var order Order

	//Decodes the request and put the content of the body into the order
	_ = json.NewDecoder(request.Body).Decode(&order)

	//Retrieves the infos about the beer he wants to order
	pizza := FindPizzaByID(order.ID)

	numberOfPizzaWanted := order.Quantity / plateQuantity
	//If the customer sends enough money
	//float32() converts a int into a float32
	if order.Money >= pizza.Price * float32(numberOfPizzaWanted) {
		plates := servePizza(&order, numberOfPizzaWanted)

		json.NewEncoder(writer).Encode(plates)
	} else {
		json.NewEncoder(writer).Encode("Not enough money")
	}
}

func BreakPlate(writer http.ResponseWriter, request *http.Request) {
	log.Println("A plate broke")
	initHeaders(writer)
	numberOfBrokenPlate++
	json.NewEncoder(writer).Encode("Tonight, ", strconv.Itoa(numberOfBrokenPlate + " plate(s) broke"))
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func servePizza(order *Order, numberOfPizzaWanted int) []plate {
	var mugs []Plate

	//We search for the good barrel
	breads, idx := FindBreadFromBakery(order.ID, breads)
	for i := 0; i < numberOfPizzaWanted; i++ {
		var plate Plate
		//When there is no more enough beer, we order a new barrel
		if (breads[idx].Quantity - plateQuantity) <= 0 {
			orderBreads(idx, bread)
		}

		plate.Pizza = breads.Pizza
		plate.Quantity = plateQuantity
		breads[idx].Quantity -= plateQuantity
		plates = append(plates, plate)
	}


	log.Println("Number of", breads.Pizza.Name ,"breads left", breads[idx].Quantity)
	return plates
}

func orderBreads(idx int, bread Bread) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(bread)
	res, _ := http.Post("http://localhost:8000/brewery", "application/json", buffer)
	var newBread Bread
	json.NewDecoder(res.Body).Decode(&newBread)
	breads[idx].Quantity += newBreads.Quantity

	log.Println("The quantity of ", bread.Pizza.Name, " bread has been refilled, it has now", breads[idx].Quantity, "breads")
}

