package concurrency

import "fmt"

const MAX_GOROUTINES = 10

func ConcurrencyProcessing() {
	users := ProcessFile()

	usersCh := make(chan []*User)
	unvisitedUsers := make(chan *User)
	go func() { usersCh <- users }()
	initializeWorkers(unvisitedUsers, usersCh, users)
	processUsers(unvisitedUsers, usersCh, len(users))
}

func initializeWorkers(unvisitedUsers <-chan *User, usersCh chan []*User, users []*User) {
	for i := 0; i < MAX_GOROUTINES; i++ {
		go func() {
			for user := range unvisitedUsers {
				SendSmsNotification(user)
				go func(user *User) {
					friendIds := user.FriendIds
					friends := []*User{}
					for _, friendId := range friendIds {
						friend, err := FindUserById(friendId, users)
						if err != nil {
							fmt.Printf("Err %v\n", err)
							continue
						}
						friends = append(friends, friend)
					}
					_, ok := <-usersCh
					if ok {
						usersCh <- friends
					}
				}(user)
			}
		}()
	}
}

func processUsers(unvisitedUsers chan<- *User, usersCh chan []*User, size int) {
	visitedUsers := make(map[string]bool)
	count := 0
	for users := range usersCh {
		for _, user := range users {
			if !visitedUsers[user.Id] {
				visitedUsers[user.Id] = true
				count++
				if count >= size {
					close(usersCh)
				}
				unvisitedUsers <- user
			}
		}
	}
}
