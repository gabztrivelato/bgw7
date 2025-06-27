package handler

import (
	output "app/internal/plataform"
	"app/internal/service"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewHandlerTicketDefault(sv *service.ServiceTicketDefault) *HandlerTicketDefault {
	return &HandlerTicketDefault{sv: sv}
}

type HandlerTicketDefault struct {
	sv *service.ServiceTicketDefault
}

func (h *HandlerTicketDefault) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			code := http.StatusInternalServerError
			body := &output.GetTotalTicketResponse{
				Message: err.Error(),
				Data:    0,
				Error:   true,
			}
			response.JSON(w, code, body)
			return
		}

		body := &output.GetTotalTicketResponse{
			Message: "Total of tickets:",
			Data:    total,
			Error:   false,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *HandlerTicketDefault) GetTicketByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")

		tickets, err := h.sv.GetTicketByDestinationCountry(country)
		if err != nil {
			code := http.StatusNotFound
			body := &output.GetTicketByDestinationResponse{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			}
			response.JSON(w, code, body)
			return
		}

		body := &output.GetTicketByDestinationResponse{
			Message: "Tickets:",
			Data:    tickets,
			Error:   false,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")

		percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			code := http.StatusNotFound
			body := &output.GetTicketAvearageByDestinationResponse{
				Message: err.Error(),
				Data:    0,
				Error:   true,
			}
			response.JSON(w, code, body)
		}

		body := &output.GetTicketAvearageByDestinationResponse{
			Message: "Average:",
			Data:    percentage,
			Error:   false,
		}
		response.JSON(w, http.StatusOK, body)
	}
}
