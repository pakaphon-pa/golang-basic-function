package groupby

type User struct {
	id        int
	username  string
	firstname string
}

func GroupbyFirstname(users []User) map[string][]User {
	group := make(map[string][]User)

	for _, user := range users {
		if v, exists := group[user.firstname]; exists {
			v = append(v, user)
		}
		group[user.firstname] = append(group[user.firstname], user)
	}

	return group
}

func GroupByCount(users []User) map[string]int {
	group := make(map[string]int)

	for _, user := range users {
		group[user.firstname] += 1
	}

	return group
}
