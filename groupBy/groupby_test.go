package groupby

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var users = &[]User{
	{
		id:        1,
		username:  "a",
		firstname: "aa",
	},
	{
		id:        2,
		username:  "b",
		firstname: "aa",
	}, {
		id:        3,
		username:  "c",
		firstname: "cc",
	},
	{
		id:        4,
		username:  "d",
		firstname: "dd",
	},
}

func TestGroupByUsername(t *testing.T) {
	result := GroupbyFirstname(*users)

	expected := map[string][]User{
		"aa": {
			{
				id:        1,
				username:  "a",
				firstname: "aa",
			},
			{
				id:        2,
				username:  "b",
				firstname: "aa",
			},
		},
		"cc": {
			{
				id:        3,
				username:  "c",
				firstname: "cc",
			},
		},
		"dd": {
			{
				id:        4,
				username:  "d",
				firstname: "dd",
			},
		},
	}

	assert.Equal(t, expected, result)
}

func TestGroupByCount(t *testing.T) {
	result := GroupByCount(*users)

	expected := map[string]int{
		"aa": 2,
		"cc": 1,
		"dd": 1,
	}

	assert.Equal(t, expected, result)
}
