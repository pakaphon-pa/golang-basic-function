package array

func Filter[T any](items []T, fn func(item T) bool) []T {
	filteredItems := []T{}
	for _, value := range items {
		if fn(value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}

func FilterItem(list []Product) []Product {
	return Filter(list, func(item Product) bool {
		return item.name == "PC"
	})
}
