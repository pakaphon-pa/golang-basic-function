package string

type Product struct {
	name string
}

func ConCatStringByArrayObject(list []Product, separator string) string {
	var result string
	for _, v := range list {
		if result == "" {
			result = v.name
		} else {
			result = result + separator + v.name
		}
	}

	return result
}
