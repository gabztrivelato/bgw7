package repositories

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gabztrivelato/bgw7/desafio-go-bases/models"
	"github.com/gabztrivelato/bgw7/desafio-go-bases/utils"
)

type TicketRepository interface {
	GetTotalTickets(destination string) (int, error)
	GetCountByPeriod(time string) (int, error)
	GetAverageDestination(destination string, total int) (float64, error)
	loadFromCsv(path string) error
}

type ticketRepository struct {
	tickets []models.Ticket
}

func (r *ticketRepository) loadFromCsv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	defer file.Close()

	var tickets []models.Ticket
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		price, err := strconv.Atoi(record[5])
		if err != nil {
			return err
		}
		ticket := models.Ticket{
			Id:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			FlightHour:  record[4],
			Price:       price,
		}
		tickets = append(tickets, ticket)
	}
	r.tickets = tickets
	return nil
}

func (r *ticketRepository) GetTotalTickets(destination string) (int, error) {
	total := 0
	for _, ticket := range r.tickets {
		if strings.EqualFold(ticket.Destination, destination) {
			total++
		}
	}
	return total, nil
}

func (r *ticketRepository) GetCountByPeriod(time string) (int, error) {
	count := 0
	for _, ticket := range r.tickets {
		hour, err := utils.GetHour(ticket.FlightHour)
		if err != nil {
			fmt.Printf("error parsing string to int: %s", ticket.FlightHour)
		}

		switch time {
		case "ínicio da manhã":
			if hour == 0 || hour <= 6 {
				count++
			}
		case "manhã":
			if hour > 6 && hour <= 12 {
				count++
			}
		case "tarde":
			if hour > 12 && hour <= 19 {
				count++
			}
		case "noite":
			if hour > 19 && hour <= 23 {
				count++
			}
		}
	}
	return count, nil
}

func (r *ticketRepository) GetAverageDestination(destination string, hour int) (float64, error) {
	var traveledInSameTime int64
	var traveledToSameDestination int64

	for _, ticket := range r.tickets {
		ticketHour, err := utils.GetHour(ticket.FlightHour)

		if err != nil {
			fmt.Printf("error parsing string to int: %s", ticket.FlightHour)
		}

		if ticketHour == hour {
			traveledInSameTime++

			if strings.EqualFold(ticket.Destination, destination) {
				traveledToSameDestination++
			}
		}
	}

	percentage := (float64(traveledToSameDestination) / float64(traveledInSameTime)) * 100
	return percentage, nil

}

func NewTicketRepository(path string) TicketRepository {
	repo := &ticketRepository{
		tickets: []models.Ticket{},
	}
	err := repo.loadFromCsv(path)
	if err != nil {
		panic(err)
	}
	return repo
}
