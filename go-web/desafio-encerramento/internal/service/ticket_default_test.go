package service_test

import (
	"app/internal/repository"
	"app/internal/service"
	"app/loader"

	"testing"

	"github.com/stretchr/testify/require"
)

const testPath = "../../loader/tickets.csv"

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		loader := loader.NewLoaderTicketCSV(testPath)
		ticket, _ := loader.Load()
		lastId := len(ticket)

		rp := repository.NewRepositoryTicketMap(ticket, lastId)
		sv := service.NewServiceTicketDefault(*rp)
		total, err := sv.GetTotalAmountTickets()

		expectedTotal := 1000
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

// Tests for TestServiceTicketDefault.GetTicketByDestinationCountry
func TestServiceTicketDefault_GetTicketByDestinationCountry(t *testing.T) {
	t.Run("success to get by destination", func(t *testing.T) {
		loader := loader.NewLoaderTicketCSV(testPath)
		ticket, _ := loader.Load()
		lastId := len(ticket)

		rp := repository.NewRepositoryTicketMap(ticket, lastId)
		sv := service.NewServiceTicketDefault(*rp)
		total, err := sv.GetTicketByDestinationCountry("Brazil")

		expectedTotal := 45
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

// Tests for TestServiceTicketDefault.GetPercentageTicketsByDestinationCountry
func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get percentage by destination country", func(t *testing.T) {
		loader := loader.NewLoaderTicketCSV(testPath)
		ticket, _ := loader.Load()
		lastId := len(ticket)

		rp := repository.NewRepositoryTicketMap(ticket, lastId)
		sv := service.NewServiceTicketDefault(*rp)
		total, err := sv.GetPercentageTicketsByDestinationCountry("Brazil")

		expectedTotal := 4.5
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
