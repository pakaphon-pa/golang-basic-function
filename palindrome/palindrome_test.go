package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeString(t *testing.T) {
	result := CheckPalindrome("10001")

	assert.Equal(t, true, result)
}
