package goroutine

import (
	"fmt"
	"time"
)

type Tag2 struct {
	Name, Type string
}

type Settings struct {
	NotificationsEnabled bool
}

type User2 struct {
	Id, Name, LastName, Status string
	Tags                       []*Tag2
	*Settings
}

type NotificationsService struct {
}

func MutiGoroutine() {
	usersToUpdate := make(chan []*User2)
	usersToNotify := make(chan *User2)

	newUsers := []*User2{
		{Name: "John", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Carl", Status: "active", Settings: &Settings{NotificationsEnabled: false}},
		{Name: "Paul", Status: "deactive", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Sam", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
	}

	existingUsers := []*User2{
		{Name: "Jessica", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Eric", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Laura", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
	}

	go filterNewUsersByStatus(usersToUpdate, newUsers)
	go updateUsers(usersToUpdate, usersToNotify, existingUsers)
	notifyUsers(usersToNotify, existingUsers)
}

func filterNewUsersByStatus(usersToUpdate chan<- []*User2, users []*User2) {
	defer close(usersToUpdate) // close channel after finish at all

	filteredUsers := []*User2{}

	for _, user := range users {
		if user.Status == "active" && user.Settings.NotificationsEnabled {
			filteredUsers = append(filteredUsers, user)
		}
	}

	usersToUpdate <- filteredUsers
}

func updateUsers(usersToUpdate <-chan []*User2, userToNotify chan<- *User2, users []*User2) {
	defer close(userToNotify) // close channel after finish at all

	for _, user := range users {
		user.Tags = append(user.Tags, &Tag2{Name: "userNotified", Type: "Notifications"})
	}

	newUsers := <-usersToUpdate

	for _, user := range newUsers {
		time.Sleep(1 * time.Second)
		user.Tags = append(user.Tags, &Tag2{Name: "NewNotification", Type: "Notifications"})
		userToNotify <- user
	}

}

func notifyUsers(userToNotify <-chan *User2, users []*User2) {
	service := &NotificationsService{}
	for _, user := range users {
		service.SendEmailNoti(user, "Tags", "A new tag has been added to your profile")
	}

	for user := range userToNotify {
		service.SendEmailNoti(user, "Tags", "You got your first tag!!")
	}
}

func (n *NotificationsService) SendEmailNoti(user *User2, title, message string) {
	fmt.Printf("Email Notification Sent to %v, Hi %s, %s\n", user, user.Name, message)
}
