package questions

import (
	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"io"
	"net/http"
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
