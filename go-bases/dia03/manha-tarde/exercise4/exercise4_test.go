package exercise4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationMinimum(t *testing.T) {
	values := []float64{7.5, 8.0, 6.0, 9.2, 5.5}
	opFunc, err := operation("minimum")
	assert.NoError(t, err)
	expected := 5.5
	result := opFunc(values...)
	assert.Equal(t, expected, result, "O valor mínimo deve ser 5.5")
}

func TestOperationMaximum(t *testing.T) {
	values := []float64{7.5, 8.0, 6.0, 9.2, 5.5}
	opFunc, err := operation("maximum")
	assert.NoError(t, err)
	expected := 9.2
	result := opFunc(values...)
	assert.Equal(t, expected, result, "O valor máximo deve ser 9.2")
}

func TestOperationAverage(t *testing.T) {
	values := []float64{7.5, 8.0, 6.0, 9.2, 5.5}
	opFunc, err := operation("average")
	assert.NoError(t, err)
	expected := 7.24
	result := opFunc(values...)
	assert.InDelta(t, expected, result, 0.01, "O valor médio deve ser aproximadamente 7.24")
}
