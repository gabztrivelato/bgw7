package repositories

import (
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	repository := NewTicketRepository("../tickets_test.csv")

	destination := "Brazil"
	expected := 1
	result, err := repository.GetTotalTickets(destination)

	if err != nil {
		t.Errorf("error calling GetTotalTickets: %v", err)
	}
	if result != expected {
		t.Errorf("expected %d ticket and got: %d", expected, result)
	}
}

func TestGetCountByPeriod(t *testing.T) {
	repository := NewTicketRepository("../tickets_test.csv")

	time := "noite"
	expected := 3
	result, err := repository.GetCountByPeriod(time)

	if err != nil {
		t.Errorf("error calling GetCountByPeriod: %v", err)
	}
	if result != expected {
		t.Errorf("expected %v and got: %v", expected, result)
	}
}

func TestGetAverageDestination(t *testing.T) {
	repository := NewTicketRepository("../tickets_test.csv")

	destination, hour := "China", 20
	expected := 100.0
	result, err := repository.GetAverageDestination(destination, hour)

	if err != nil {
		t.Errorf("error calling GetAverageDestination: %v", err)
	}
	if result != expected {
		t.Errorf("expected %v and got : %v", expected, result)
	}

}
