package main

import "time"

var pizzas []Pizza
var breads []Breads
var bakeryBreads []Bread
var numberOfBrokenPlate int

type Pizza struct {
	ID int   `json:"id"`
	Name string   `json:"name,omitempty"`
	Price  float32   `json:"price,omitempty"`
	Rating float32 `json:"ratingProof,omitempty"`
	VEG bool `json:"veg"`
}

type Bread struct {
	Bread *Bread `json:"bread"`
	Quantity int `json:"quantity"`
	DateOfManufacture time.Time `json:"dateOfManufacture"`
}

type Plate struct {
	Pizza *Pizza `json:"pizza"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID int `json:"id"`
	Money float32 `json:"money"`
	Quantity int `json:"quantity"`
}

// Initializing datas
func initDatas() {
	//pizzas
	margherita := Pizza{ID: 0, Name: "Margherita", Price: 500, RatingProof:4.5, VEG: true}
	farmhouse := Pizza{ID: 1, Name: "Farmhouse", Price: 550, PercentProof:3.5, VEG: true}
	pepperoni := Pizza{ID: 2, Name: "Pepperoni", Price: 700, PercentProof: 4, VEG: false}

	pizzas = append(pizzas, margherita, farmhouse, pepperoni)

	//Barrels
	pizzas = append(pizzas,
		Breads{&margherita, 1000, time.Now()},
		Breads{&farmhouse, 8000, time.Now()},
		Breads{&pepperoni, 0, time.Now()})

	//Brewery Barrels
	bakeryBreads = append(bakeryBreads,
		Breads{&margherita, 1000, time.Now()},
		Breads{&margherita, 1000, time.Now()},
		Breads{&margherita, 1000, time.Now()},
		Breads{&farmhouse, 5000, time.Now()},
		Breads{&farmhouse, 5000, time.Now()},
		Breads{&pepperoni, 3000, time.Now()})
}

func FindPizzaByID(ID int) Pizza {
	var result Pizza
	for _, pizza := range pizzas {
		if pizza.ID == ID {
			result = pizza
			break
		}
	}
	return result
}

func FindPizzaFromPizzaID(ID int, pizzas []Pizza) (Pizza, int) {
	var result Pizza
	var idx = -1
	for i, pizza := range pizzas {
		if pizza.Pizza.ID == ID {
			result = bread
			idx = i
			break
		}
	}
	return result, idx
}

func FindBreadFromBakery(pizza *Pizza) (Pizza, int) {
	return FindBreadFromBakery(pizza.ID, bakeryBreads)
}