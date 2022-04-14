package romanNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	result := RomanToInt("III")

	assert.Equal(t, 3, result)
}

func TestRomanToInt2(t *testing.T) {
	result := RomanToInt("CCC")

	assert.Equal(t, 300, result)
}

func TestIntToRoman(t *testing.T) {
	result := IntToRoman(3)

	assert.Equal(t, "III", result)
}

func TestIntToRoman2(t *testing.T) {
	result := IntToRoman(300)

	assert.Equal(t, "CCC", result)
}
