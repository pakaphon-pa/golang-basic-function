package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Users = []User{
	{
		"B", "C",
	},
	{
		"A", "A",
	},
	{
		"B", "A",
	},
	{
		"A", "Z",
	},
}

func TestSortDataList(t *testing.T) {
	result := SortData(DataList)
	expected := []Product{
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
		{
			name:  "PC",
			price: 50000.00,
			star:  3,
		},
	}
	assert.Equal(t, expected, result)
}

func TestSortMultiKey(t *testing.T) {
	result := SortMultikey(Users)

	expected := []User{
		{
			"A", "A",
		},
		{
			"A", "Z",
		},
		{
			"B", "A",
		},
		{
			"B", "C",
		},
	}
	assert.Equal(t, expected, result)
}
