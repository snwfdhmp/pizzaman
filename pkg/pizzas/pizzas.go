package pizzas

import (
	"encoding/json"
	"errors"
	"github.com/snwfdhmp/pizzaman/pkg/fridge"
	"io"
)

type Pizza struct {
	Sauce  string `json:"sauce,omitempty"`
	Meat   string `json:"meat,omitempty"`
	Cheese string `json:"cheese,omitempty"`
}

func (p *Pizza) addSauce(s string) error {
	if !fridge.HasSauce(s) {
		return errors.New("this sauce is not available")
	}
	p.Sauce = s
	return nil
}

func (p *Pizza) addMeat(s string) error {
	if !fridge.HasMeat(s) {
		return errors.New("this meat is not available")
	}
	p.Meat = s
	return nil
}

func (p *Pizza) addCheese(s string) error {
	if !fridge.HasCheese(s) {
		return errors.New("this cheese is not available")
	}
	p.Cheese = s
	return nil
}

func (p *Pizza) Write(w io.Writer) {
	b, _ := json.Marshal(p)
	w.Write(b)
}

func New(sauce, meat, cheese string) (Pizza, error) {
	p := Pizza{}

	err := p.addSauce(sauce)
	if err != nil {
		return Pizza{}, err
	}

	err = p.addMeat(meat)
	if err != nil {
		return Pizza{}, err
	}

	err = p.addCheese(cheese)
	if err != nil {
		return Pizza{}, err
	}

	return p, nil
}
