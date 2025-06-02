package exercise2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAverage(t *testing.T) {
	var grades = []int64{7, 8, 9, 6, 8}
	average := calcAverage(grades...)
	expectedAverage := 7.6

	assert.Equal(t, expectedAverage, average, "The average should be 8.0")
}
