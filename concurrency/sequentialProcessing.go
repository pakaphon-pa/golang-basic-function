package concurrency

import "fmt"

func SequentialProcessing() {
	users := ProcessFile()
	visited := make(map[string]bool)

	for _, user := range users {
		if !visited[user.Id] {
			visited[user.Id] = true
			SendSmsNotification(user)
			for _, friendId := range user.FriendIds {
				friend, err := FindUserById(friendId, users)
				if err != nil {
					fmt.Printf("Err %v\n", err)
					continue
				}

				if !visited[friend.Id] {
					visited[friend.Id] = true
					SendSmsNotification(friend)
				}
			}
		}
	}

}
