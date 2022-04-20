package array

func RemoveDuplicate(slice []string) []string {
	dic := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := dic[entry]; !value {
			dic[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
