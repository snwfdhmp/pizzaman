package fridge

var (
	Sauces  []string
	Meats   []string
	Cheeses []string
)

func HasSauce(s string) bool {
	for _, sauce := range Sauces {
		if sauce == s {
			return true
		}
	}
	return false
}

func HasMeat(s string) bool {
	for _, meat := range Meats {
		if meat == s {
			return true
		}
	}
	return false
}

func HasCheese(s string) bool {
	for _, cheese := range Cheeses {
		if cheese == s {
			return true
		}
	}
	return false
}

func FillSauces(s []string) {
	for _, sauce := range s {
		Sauces = append(Sauces, sauce)
	}
}

func FillMeats(s []string) {
	for _, meat := range s {
		Meats = append(Meats, meat)
	}
}

func FillCheeses(s []string) {
	for _, cheese := range s {
		Cheeses = append(Cheeses, cheese)
	}
}
