package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {

		// given
		whiteShark := hunt.NewWhiteShark(true, false, 30.0)
		tuna := hunt.NewTuna("Little Tuna", 15.0)
		expectedHungry := false
		expectedTired := true

		// when
		result := whiteShark.Hunt(tuna)

		//then
		assert.Nil(t, result)
		assert.Equal(t, expectedHungry, whiteShark.Hungry)
		assert.Equal(t, expectedTired, whiteShark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// given
		whiteShark := hunt.NewWhiteShark(false, false, 30.0)
		tuna := hunt.NewTuna("Litte Tuna", 15.0)
		expectedError := hunt.ErrSharkIsNotHungry

		// when
		result := whiteShark.Hunt(tuna)

		// then
		assert.Error(t, result)
		assert.Equal(t, expectedError, result)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// given
		whiteShark := hunt.NewWhiteShark(true, true, 30.0)
		tuna := hunt.NewTuna("Little Tuna", 15.0)
		expectedError := hunt.ErrSharkIsTired

		// when
		result := whiteShark.Hunt(tuna)

		// then
		assert.Error(t, result)
		assert.Equal(t, expectedError, result)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// given
		whiteShark := hunt.NewWhiteShark(true, false, 15.0)
		tuna := hunt.NewTuna("Speed Tuna", 25.0)
		expectedError := hunt.ErrSharkIsSlower

		// when
		result := whiteShark.Hunt(tuna)

		// then
		assert.Error(t, result)
		assert.Equal(t, expectedError, result)

	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// given
		whiteShark := hunt.NewWhiteShark(true, false, 30.0)
		tuna := (*hunt.Tuna)(nil)
		expectedError := hunt.ErrTunaIsNil

		// when
		result := whiteShark.Hunt(tuna)

		// then
		assert.Error(t, result)
		assert.Equal(t, expectedError, result)
	})
}
