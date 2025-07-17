package prey_test

import (
	"testdoubles/positioner"
	prey "testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeedDefault(t *testing.T) {
	tuna := prey.CreateTuna()
	result := tuna.GetSpeed()
	assert.Equal(t, tuna.GetSpeed(), result)
}

func TestGetSpeedWithValue(t *testing.T) {
	position := &positioner.Position{X: 10, Y: 20, Z: 5}
	tuna := prey.NewTuna(100, position)
	expectedSpeed := 100.0

	result := tuna.GetSpeed()

	assert.Equal(t, expectedSpeed, result)
}

func TestGetPositionDefault(t *testing.T) {
	tuna := prey.CreateTuna()
	result := tuna.GetPosition()
	assert.Equal(t, tuna.GetPosition(), result)
}

func TestGetPositionWithValue(t *testing.T) {
	position := &positioner.Position{X: 10, Y: 20, Z: 10}
	tuna := prey.NewTuna(100, position)
	expectedPosition := position

	result := tuna.GetPosition()

	assert.Equal(t, expectedPosition, result)
}
