package exercise3

import "testing"

func TestCalcSalaryWithCategoryA(t *testing.T) {
	minutesWorked := int64(120)
	category := "A"
	expectedSalary := 9000.0

	salary := calcSalary(minutesWorked, category)

	if salary != expectedSalary {
		t.Errorf("Expected salary: %f, got: %f", expectedSalary, salary)
	}
}

func TestCalcSalaryWithCategoryB(t *testing.T) {
	minutesWorked := int64(120)
	category := "B"
	expectedSalary := 3600.0

	salary := calcSalary(minutesWorked, category)

	if salary != expectedSalary {
		t.Errorf("Expected salary: %f, got: %f", expectedSalary, salary)
	}
}

func TestCalcSalaryWithCategoryC(t *testing.T) {
	minutesWorked := int64(120)
	category := "C"
	expectedSalary := 2000.0

	salary := calcSalary(minutesWorked, category)

	if salary != expectedSalary {
		t.Errorf("Expected salary: %f, got: %f", expectedSalary, salary)
	}
}
