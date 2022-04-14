package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataList = []Product{
	{
		name:  "PC",
		price: 50000.00,
		start: 3,
	},
	{
		name:  "PS5",
		price: 16000.00,
		start: 2,
	},
	{
		name:  "Xbox",
		price: 16000.00,
		start: 2,
	},
}

func TestMaxPriceInProductList(t *testing.T) {
	result := FindMaxPriceProduct(dataList)
	expected := Product{
		name:  "PC",
		price: 50000.00,
		start: 3,
	}
	assert.Equal(t, expected, result)
}
