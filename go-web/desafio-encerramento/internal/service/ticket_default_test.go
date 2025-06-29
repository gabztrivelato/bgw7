package service_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		// sv := service.NewServiceTicketDefault(rp)

		// // act
		// total, err := sv.GetTotalTickets()

		// // assert
		// expectedTotal := 1
		// require.NoError(t, err)
		// require.Equal(t, expectedTotal, total)
	})
}
