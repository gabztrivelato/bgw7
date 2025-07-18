package handler

import (
	"encoding/json"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"

	"github.com/bootcamp-go/web/response"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigPrey
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// process
		prey := prey.NewTuna(requestBody.Speed, requestBody.Position)
		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "Prey configured successfully", "data": prey})
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigHunter
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// process
		simulatorCreated := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{
			MaxTimeToCatch: 10.0,
			Positioner:     positioner.NewPositionerDefault(),
		})
		configHunter := &hunter.ConfigWhiteShark{
			Speed:     requestBody.Speed,
			Position:  requestBody.Position,
			Simulator: simulatorCreated,
		}
		hunterCreated := hunter.NewWhiteShark(*configHunter)
		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "Hunter configured successfully", "data": hunterCreated})
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// process

		// response
	}
}
