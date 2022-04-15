package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInProductList(t *testing.T) {
	result := FilterItem(DataList)
	expected := []Product{
		{
			name:  "PC",
			price: 50000.00,
			star:  3,
		},
	}
	assert.Equal(t, expected, result)
}
