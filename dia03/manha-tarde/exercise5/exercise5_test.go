package exercise5

import "testing"

func TestDog(t *testing.T) {
	quantity := 10
	expected := 100.0
	result := animalDog(quantity)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
func TestCat(t *testing.T) {
	quantity := 10
	expected := 50.0
	result := animalCat(quantity)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestHamster(t *testing.T) {
	quantity := 10
	expected := 2.5
	result := animalHamster(quantity)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestTarantula(t *testing.T) {
	quantity := 10
	expected := 1.5
	result := animalTarantula(quantity)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
