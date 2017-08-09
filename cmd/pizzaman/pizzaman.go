package main

import (
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"github.com/snwfdhmp/pizzaman/pkg/orders"
	"github.com/snwfdhmp/pizzaman/pkg/routes"
	"log"
	"net/http"
)

var (
	sauces  []string
	meats   []string
	cheeses []string
)

func main() {
	log.Println("Welcome to your pizzeria")

	meats = append(meats, "chicken")

	cheeses = append(cheeses, "goat")

	log.Println("Buying sauces")
	buySauces()

	log.Println("Buying meats")
	buyMeats()

	log.Println("Buying cheeses")
	buyCheeses()

	log.Println("Filling fridge")
	fridge.FillSauces(sauces)
	fridge.FillMeats(meats)
	fridge.FillCheeses(cheeses)

	log.Println("Switching on phone")
	router := routes.Init()

	log.Println("Waking up employees")
	go orders.PrepareAll()

	log.Fatal(http.ListenAndServe(":9876", router))
}

func buySauces() {
	sauces = append(sauces, "tomato")
	sauces = append(sauces, "ketchup")
	sauces = append(sauces, "classico")
	sauces = append(sauces, "italian")
	sauces = append(sauces, "buffalo")
	sauces = append(sauces, "supreme")
	sauces = append(sauces, "no")
}

func buyMeats() {
	meats = append(meats, "chicken")
	meats = append(meats, "bacon")
	meats = append(meats, "no")
}

func buyCheeses() {
	cheeses = append(cheeses, "goat")
	cheeses = append(cheeses, "mozzarella")
	cheeses = append(cheeses, "emmental")
	cheeses = append(cheeses, "ricotta")
	cheeses = append(cheeses, "provolone")
	cheeses = append(cheeses, "cheddar")
	cheeses = append(cheeses, "romano")
	cheeses = append(cheeses, "no")
}
