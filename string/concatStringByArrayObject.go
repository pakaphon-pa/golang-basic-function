package string

type Product struct {
	name string
}

func ConCatStringProductname(list []Product, separator string) string {
	var result string
	for _, v := range list {
		result = ConCatStringByArrayObject(result, v.name, separator)
	}
	return result
}

func ConCatStringByArrayObject(string1 string, string2 string, separator string) string {
	result := string1
	if result == "" {
		result = string2
	} else {
		result = result + separator + string2
	}
	return result
}
