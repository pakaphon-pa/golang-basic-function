package array

import (
	"sort"
)

type User struct {
	Firstname string
	Lastname  string
}

func SortData(list []Product) []Product {
	sort.Slice(list, func(i, j int) bool {
		return list[i].star < list[j].star
	})

	return list
}

func SortMultikey(list []User) []User {
	sort.SliceStable(list, func(i, j int) bool {
		if list[i].Firstname != list[j].Firstname {
			return list[i].Firstname < list[j].Firstname
		}

		return list[i].Lastname < list[j].Lastname
	})

	return list
}
