package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceSumPrice(t *testing.T) {
	result := Sum(DataList)
	expected := 82000
	assert.Equal(t, expected, result)
}
