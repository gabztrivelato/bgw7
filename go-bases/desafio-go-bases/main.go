package main

import (
	"fmt"

	"github.com/gabztrivelato/bgw7/desafio-go-bases/repositories"
	"github.com/gabztrivelato/bgw7/desafio-go-bases/services"
)

func main() {
	repository := repositories.NewTicketRepository("tickets.csv")
	service := services.NewTicketService(repository)

	averageDestination, err := service.GetAverageDestination("China", 20)
	if err != nil {
		panic(err)
	}

	countPeriod, err := service.GetCountByPeriod("tarde")
	if err != nil {
		panic(err)
	}

	totalTickets, err := service.GetTotalTickets("Brazil")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Average destination: %+v\n", averageDestination)
	fmt.Printf("Ticket per period: %+v\n", countPeriod)
	fmt.Printf("Total tickets:: %+v\n", totalTickets)

}
