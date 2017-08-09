package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/orders"
)

func Init() *httprouter.Router {
	r := httprouter.New()

	r.GET("/order/:sauce/:meat/:cheese", orders.Take)

	return r
}
