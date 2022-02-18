package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main(){
	fmt.Println("Welcome to the The Pizza Hut !")
	router := mux.NewRouter()

	initDatas()
	buildPizzaHutRoutes(router)
	buildBakeryRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Bar's routes
func buildPizzaHutRoutes(router *mux.Router) {
	prefix := "/PizzaHut"
	router.HandleFunc(prefix, GetInfo).Methods("GET")
	router.HandleFunc(prefix + "/{id}", GetPizzaInfo).Methods("GET")
	router.HandleFunc(prefix, OrderPizza).Methods("POST")
	router.HandleFunc(prefix, BreakPlate).Methods("DELETE")
}

//Brewery's routes
func buildBreweryRoutes(router *mux.Router) {
	prefix := "/bakery"
	router.HandleFunc(prefix, OrderBreads).Methods("POST")
	router.HandleFunc(prefix, ProduceBreads).Methods("PUT")
}


