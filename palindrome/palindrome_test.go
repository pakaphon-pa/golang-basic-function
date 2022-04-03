package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeString(t *testing.T) {
	result := CheckPalindrome("ABABA")

	assert.Equal(t, true, result)
}

func TestIsPalindromeStringFalse(t *testing.T) {
	result := CheckPalindrome("TIME")

	assert.Equal(t, false, result)
}

func TestIsPalindromeInt(t *testing.T) {
	result := CheckPalindrome(10001)

	assert.Equal(t, true, result)
}
