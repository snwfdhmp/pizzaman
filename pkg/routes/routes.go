package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/orders"
	"github.com/snwfdhmp/pizzaman/pkg/questions"
)

func Init() *httprouter.Router {
	r := httprouter.New()

	r.GET("/order/:sauce/:meat/:cheese", orders.Take)

	//r.GET("/whatsaucesdoyouhave", fridge.WhatSauces)
	//r.GET("/whatmeatsdoyouhave", fridge.WhatMeats)
	r.GET("/whatcheesesdoyouhave", questions.WhatCheeses)

	return r
}
