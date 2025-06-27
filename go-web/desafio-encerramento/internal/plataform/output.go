package output

import "app/internal"

type GetTotalTicketResponse struct {
	Message string
	Data    int
	Error   bool
}

type GetTicketByDestinationResponse struct {
	Message string
	Data    map[int]internal.TicketAttributes
	Error   bool
}

type GetTicketAvearageByDestinationResponse struct {
	Message string
	Data    float64
	Error   bool
}
