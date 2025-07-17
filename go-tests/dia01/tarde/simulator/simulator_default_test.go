package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCatch(t *testing.T) {
	t.Run("case 1: Hunter gets the prey", func(t *testing.T) {
		positioner_default := positioner.NewPositionerDefault()
		simulator_default := simulator.NewCatchSimulatorDefault(100.0, positioner_default)
		hunter := &simulator.Subject{&positioner.Position{X: 10, Y: 40, Z: 5}, 100.0}
		prey := &simulator.Subject{&positioner.Position{X: 10, Y: 50, Z: 5}, 50.0}
		expectedCatch := true

		result := simulator_default.CanCatch(hunter, prey)
		assert.Equal(t, expectedCatch, result)
	})

	t.Run("case 2: Hunter doesn't get the prey because is slower", func(t *testing.T) {
		positioner_default := positioner.NewPositionerDefault()
		simulator_default := simulator.NewCatchSimulatorDefault(100.0, positioner_default)
		hunter := &simulator.Subject{&positioner.Position{X: 10, Y: 40, Z: 5}, 40.0}
		prey := &simulator.Subject{&positioner.Position{X: 10, Y: 50, Z: 5}, 50.0}
		expectedCatch := false

		result := simulator_default.CanCatch(hunter, prey)
		assert.Equal(t, expectedCatch, result)
	})

	t.Run("case 3: Hunter doesn't get the prey because time isn't enough", func(t *testing.T) {
		positioner_default := positioner.NewPositionerDefault()
		simulator_default := simulator.NewCatchSimulatorDefault(10.0, positioner_default)
		hunter := &simulator.Subject{&positioner.Position{X: 10, Y: 40, Z: 5}, 55.0}
		prey := &simulator.Subject{&positioner.Position{X: 50, Y: 130, Z: 10}, 50.0}
		expectedCatch := false

		result := simulator_default.CanCatch(hunter, prey)
		assert.Equal(t, expectedCatch, result)
	})
}
