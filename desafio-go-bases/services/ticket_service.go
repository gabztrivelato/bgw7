package services

import (
	"github.com/gabztrivelato/bgw7/desafio-go-bases/repositories"
)

type TicketService struct {
	ticketRepository repositories.TicketRepository
}

func NewTicketService(r repositories.TicketRepository) *TicketService {
	return &TicketService{ticketRepository: r}
}

func (s *TicketService) GetTotalTickets(destination string) (int, error) {
	return s.ticketRepository.GetTotalTickets(destination)
}

func (s *TicketService) GetCountByPeriod(time string) (int, error) {
	return s.ticketRepository.GetCountByPeriod(time)
}

func (s *TicketService) GetAverageDestination(destination string, hour int) (float64, error) {
	return s.ticketRepository.GetAverageDestination(destination, hour)
}
