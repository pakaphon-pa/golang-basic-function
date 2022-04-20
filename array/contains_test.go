package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var List = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

func TestContainsData(t *testing.T) {

	result := Contains(List, "A")
	expected := true
	assert.Equal(t, expected, result)
}
