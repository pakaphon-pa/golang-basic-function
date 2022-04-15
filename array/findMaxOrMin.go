package array

type Product struct {
	name  string
	price float32
	star  int
}

func FindMaxPriceProduct(list []Product) Product {
	result := list[0]

	for _, value := range list {
		if value.price > result.price {
			result = value
		}
	}

	return result
}

func FindMinPriceProduct(list []Product) Product {
	result := list[0]

	for _, value := range list {
		if value.price < result.price {
			result = value
		}
	}

	return result
}
