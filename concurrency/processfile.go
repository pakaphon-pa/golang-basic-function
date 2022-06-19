package concurrency

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var count = 1

func ProcessFile() []*User {
	f, err := os.Open("concurrency/students2.csv")

	if err != nil {
		log.Fatal(err)
	}

	users := scanFile(f)

	return users
}

func SendSmsNotification(user *User) {
	time.Sleep(10 * time.Microsecond)

	fmt.Printf("Sending sms notification to %v %v\n", user.Phone, count)
	count++
}

func FindUserById(userId string, users []*User) (*User, error) {
	for _, user := range users {
		if user.Id == userId {
			return user, nil
		}
	}

	return nil, fmt.Errorf("User not found with id %v", userId)
}

func scanFile(f *os.File) []*User {
	s := bufio.NewScanner(f)
	users := []*User{}

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		ids := strings.Split(lineArray[5], " ")
		ids = ids[1 : len(ids)-1]

		user := &User{
			Id:        lineArray[0],
			Name:      lineArray[1],
			LastName:  lineArray[2],
			Email:     lineArray[3],
			Phone:     lineArray[4],
			FriendIds: ids,
		}
		users = append(users, user)
	}

	return users
}
