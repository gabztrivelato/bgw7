package repository

import (
	"app/internal"
	"context"
	"errors"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}
	if len(t) == 0 {
		return t, errors.New("tickets not found")
	}

	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by country
func (r *RepositoryTicketMap) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	ticketsByDestination, err := r.GetTicketsByDestinationCountry(country)
	if err != nil {
		return 0, err
	}
	percentage = (float64(len(ticketsByDestination)) / float64(len(r.db))) * 100
	return percentage, nil
}

// GetTotalAmountTickets returns the amount of tickets
func (r *RepositoryTicketMap) GetTotalAmountTickets() (total int, err error) {
	total = len(r.db)
	if total == 0 {
		return 0, errors.New("tickets not found")
	}
	return
}
