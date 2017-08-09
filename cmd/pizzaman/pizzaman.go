package main

import (
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"github.com/snwfdhmp/pizzaman/pkg/orders"
	"github.com/snwfdhmp/pizzaman/pkg/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Welcome to your pizzeria")

	var sauces []string
	var meats []string
	var cheeses []string

	sauces = append(sauces, "tomato")
	meats = append(meats, "chicken")
	cheeses = append(cheeses, "goat")

	fridge.FillSauces(sauces)
	fridge.FillMeats(meats)
	fridge.FillCheeses(cheeses)

	log.Println("Fridge filled")

	router := routes.Init()
	log.Println("Phone switched on")

	go orders.PrepareAll()
	log.Println("Ready to prepare orders")

	log.Fatal(http.ListenAndServe(":9876", router))
}
