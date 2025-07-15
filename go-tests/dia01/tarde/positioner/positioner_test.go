package positioner_test

import (
	"math"
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("case 1: all coordinates are negative", func(t *testing.T) {
		// given
		pos1 := &positioner.Position{X: -1.0, Y: -2.0, Z: -3.0}
		pos2 := &positioner.Position{X: -4.0, Y: -6.0, Z: -3.0}
		positioner := positioner.NewPositionerDefault()
		expectedResult := math.Sqrt(math.Pow(-1.0-(-4.0), 2.0) + math.Pow(-2.0-(-6.0), 2.0) + math.Pow(-3.0-(-3.0), 2.0))

		//when
		result := positioner.GetLinearDistance(pos1, pos2)

		// then
		assert.Equal(t, expectedResult, result)
	})

	t.Run("case 2: all cordinates are positve", func(t *testing.T) {
		// given
		pos1 := &positioner.Position{X: 1.0, Y: 2.0, Z: 3.0}
		pos2 := &positioner.Position{X: 4.0, Y: 6.0, Z: 3.0}
		positioner := positioner.NewPositionerDefault()
		expectedResult := math.Sqrt(math.Pow(1.0-(4.0), 2.0) + math.Pow(2.0-(6.0), 2.0) + math.Pow(3.0-(3.0), 2.0))

		// when
		result := positioner.GetLinearDistance(pos1, pos2)

		//then
		assert.Equal(t, expectedResult, result)
	})

	t.Run("case 3: cordinates return linear distante without decimal", func(t *testing.T) {
		// given
		pos1 := &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0}
		pos2 := &positioner.Position{X: 3.0, Y: 4.0, Z: 0.0}
		positioner := positioner.NewPositionerDefault()
		expectedResult := 5.0

		// when
		result := positioner.GetLinearDistance(pos1, pos2)

		// then
		assert.Equal(t, expectedResult, result)
	})
}
