package exercise1

func calcTax(income float64) float64 {
	var tax float64
	switch {
	case income > 150000.0:
		tax = 27.0
	case income > 50000.0:
		tax = 17.0
	default:
		tax = 0.0
	}
	return tax
}
