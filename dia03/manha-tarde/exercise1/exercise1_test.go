package exercise1

import "testing"

func TestCalculateTaxIncomeUnder50000(t *testing.T) {
	income := 50000.0

	expectedTax := 0.0

	var tax = calcTax(income)

	if tax != expectedTax {
		t.Errorf("Expected tax for income %.2f to be %.2f, but got %.2f", income, expectedTax, tax)
	}
}

func TestCalculateTaxIncomeOver50000(t *testing.T) {
	income := 60000.0

	expectedTax := 17.0

	var tax = calcTax(income)

	if tax != expectedTax {
		t.Errorf("Expected tax for income %.2f to be %.2f, but got %.2f", income, expectedTax, tax)
	}
}

func TestCalculateTaxIncomeOver150000(t *testing.T) {
	income := 160000.0

	expectedTax := 27.0

	var tax = calcTax(income)

	if tax != expectedTax {
		t.Errorf("Expected tax for income %.2f to be %.2f, but got %.2f", income, expectedTax, tax)
	}
}
