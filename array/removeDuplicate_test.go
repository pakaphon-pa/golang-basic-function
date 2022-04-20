package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var duplicateList = []string{"A", "A", "A", "D", "E", "F", "G", "H"}

func TestRemoveDuplicateData(t *testing.T) {

	result := RemoveDuplicate(duplicateList)
	expected := []string{"A", "D", "E", "F", "G", "H"}
	assert.Equal(t, expected, result)
}
