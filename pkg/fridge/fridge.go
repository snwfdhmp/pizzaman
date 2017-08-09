package fridge

var (
	sauces  []string
	meats   []string
	cheeses []string
)

func HasSauce(s string) bool {
	for _, sauce := range sauces {
		if sauce == s {
			return true
		}
	}
	return false
}

func HasMeat(s string) bool {
	for _, meat := range meats {
		if meat == s {
			return true
		}
	}
	return false
}

func HasCheese(s string) bool {
	for _, cheese := range cheeses {
		if cheese == s {
			return true
		}
	}
	return false
}

func FillSauces(s []string) {
	for _, sauce := range s {
		sauces = append(sauces, sauce)
	}
}

func FillMeats(s []string) {
	for _, meat := range s {
		meats = append(meats, meat)
	}
}

func FillCheeses(s []string) {
	for _, cheese := range s {
		cheeses = append(cheeses, cheese)
	}
}
