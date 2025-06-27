package service

import (
	"app/internal"
	"app/internal/repository"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicketMap
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicketMap) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	total, err = s.rp.GetTotalAmountTickets()
	return
}

// GetTicketByDestinationCountry returns the tickets filtered by country
func (s *ServiceTicketDefault) GetTicketByDestinationCountry(country string) (t map[int]internal.TicketAttributes, err error) {
	t, err = s.rp.GetTicketsByDestinationCountry(country)
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	percentage, err = s.rp.GetPercentageTicketsByDestinationCountry(country)
	return
}
