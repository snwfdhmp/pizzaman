package questions

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"github.com/snwfdhmp/pizzaman/pkg/orders"
)

func list(w io.Writer, arr []string) {
	io.WriteString(w, "We have :\n")
	for _, el := range arr {
		if el != "no" {
			io.WriteString(w, "- "+el+"\n")
		}
	}
}

func WhatSauces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list(w, fridge.Sauces)
}

func WhatMeats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list(w, fridge.Meats)
}

func WhatCheeses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list(w, fridge.Cheeses)
}

func HowIsMyPizza(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	state := orders.State(ps.ByName("ticket"))

	io.WriteString(w, state)
}
