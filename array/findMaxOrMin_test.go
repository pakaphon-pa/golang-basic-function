package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var DataList = []Product{
	{
		name:  "PC",
		price: 50000.00,
		star:  3,
	},
	{
		name:  "PS5",
		price: 16000.00,
		star:  1,
	},
	{
		name:  "Xbox",
		price: 16000.00,
		star:  2,
	},
}

func TestMaxPriceInProductList(t *testing.T) {
	result := FindMaxPriceProduct(DataList)
	expected := Product{
		name:  "PC",
		price: 50000.00,
		star:  3,
	}
	assert.Equal(t, expected, result)
}

func TestMixPriceInProductList(t *testing.T) {
	result := FindMinPriceProduct(DataList)
	expected := Product{
		name:  "PS5",
		price: 16000.00,
		star:  1,
	}
	assert.Equal(t, expected, result)
}
