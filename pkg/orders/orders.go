package orders

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/snwfdhmp/pizzaman/pkg/pizzas"
)

var (
	toBePrepared OrderChan
	orders       Orders
)

type OrderSuccessResp struct {
	Message string       `json:"message"`
	Pizza   pizzas.Pizza `json:"pizza"`
	Ticket  string       `json:"ticket"`
}

type Order struct {
	Pizza    pizzas.Pizza
	Ticket   string
	Prepared bool
}

type Orders []Order
type OrderChan chan Order

func (r OrderSuccessResp) Write(w io.Writer) {
	b, _ := json.Marshal(r)
	w.Write(b)
}

func Take(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sauce := ps.ByName("sauce")
	meat := ps.ByName("meat")
	cheese := ps.ByName("cheese")

	p, err := pizzas.New(sauce, meat, cheese)
	if err != nil {
		log.Println("Cannot create pizza:", err)
		io.WriteString(w, err.Error())
		return
	}

	t := randTicket()

	OrderSuccessResp{
		Message: "Your order has been taken",
		Pizza:   p,
		Ticket:  t,
	}.Write(w)

	order := Order{Pizza: p, Ticket: t}

	orders = append(orders, order)
	toBePrepared.Prepare(order)
}

func randTicket() string {
	b := make([]byte, 2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (o *Order) Prepare() {
	err := wait("6s")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(o.Ticket, "finished")
		MarkAsPrepared(o.Ticket)
	}
}

func (o OrderChan) Prepare(order Order) {
	o <- order
	log.Println("Command", order.Ticket, order.Pizza)
}

func PrepareAll() {
	toBePrepared = make(chan Order)
	for o := range toBePrepared {
		go func(o *Order) {
			o.Prepare()
		}(&o)
	}
}

func State(ticket string) string {
	for _, o := range orders {
		if o.Ticket != ticket {
			continue
		}

		if o.Prepared {
			return "prepared"
		}

		return "pending"
	}
	return "not found"
}

func MarkAsPrepared(ticket string) {
	for i, o := range orders {
		if o.Ticket != ticket {
			continue
		}

		orders[i].Prepared = true
	}
}

func wait(duration string) error {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return errors.New("cannot parse duration")
	}

	time.Sleep(d)

	return nil
}
