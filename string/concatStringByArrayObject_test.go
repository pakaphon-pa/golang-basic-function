package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var DataList = []Product{
	{
		name: "PC",
	},
	{
		name: "PS5",
	},
	{
		name: "Xbox",
	},
}

func TestConcatNameProduct(t *testing.T) {
	result := ConCatStringByArrayObject(DataList, ",")
	expected := "PC,PS5,Xbox"
	assert.Equal(t, expected, result)
}
