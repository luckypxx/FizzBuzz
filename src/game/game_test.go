package game

import (
	"github.com/stretchr/testify-master/assert"
	"testing"
)

var number = 68

func TestIsSpecialNumber_isSpecialNumberMultiple(t *testing.T) {
	specialNumber := 2

	assert.Equal(t, true, IsSpecialNumber(number, specialNumber, FIZZ))
	assert.Equal(t, print, FIZZ)
}

func TestIsSpecialNumber_ContainSpecialNumber(t *testing.T) {
	specialNumber := 6

	assert.Equal(t, true, IsSpecialNumber(number, specialNumber, FIZZ))
	assert.Equal(t, print, FIZZ)
}
func TestIsSpecialNumber_BothFalse(t *testing.T) {
	specialNumber := 5

	assert.Equal(t, false, IsSpecialNumber(number, specialNumber, BUZZ))
	assert.Equal(t, print, BUZZ)

}

func TestIsFizzBuzz_yes(t *testing.T) {
	firstSpecialNumber := 4
	secondSpecialNumber := 2

	assert.Equal(t, true, IsFizzBuzz(number, firstSpecialNumber, secondSpecialNumber))
	assert.Equal(t, print, FIZZ+BUZZ)
}

func TestIsFizzBuzz_no(t *testing.T) {
	firstSpecialNumber := 6
	secondSpecialNumber := 5

	assert.Equal(t, false, IsFizzBuzz(number, firstSpecialNumber, secondSpecialNumber))
	assert.Equal(t, print, FIZZ+BUZZ)
}