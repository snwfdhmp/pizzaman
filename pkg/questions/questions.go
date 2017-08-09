package questions

import (
	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"io"
	"net/http"
)

func WhatCheeses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "We have :\n")
	for _, cheese := range fridge.Cheeses {
		io.WriteString(w, "- "+cheese+"\n")
	}
}
