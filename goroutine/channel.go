package goroutine

import (
	"fmt"
	"time"
)

var (
	defaultags = []string{"SystemUser", "User", "NewUser", "System"}
)

type Tag struct {
	Name, Type string
}

type User struct {
	Id, Name, LastName, Status string
	Tags                       []*Tag
}

type Post struct {
	Title  string
	Status string
	UserId string
}

func Channel() {
	user := &User{}
	done := make(chan bool)

	go func() {
		fmt.Println("[2nd go routine] Start Building User")
		buildingUser(user)
		fmt.Println("[2nd go routine] Finished Building User")
		done <- true

		fmt.Println("[2nd go routine] Set Default user tags")
		setDefaultTags(user)
	}()

	fmt.Println("[Main go routine] Start Importing posts")
	posts := importingPosts()
	fmt.Println("[Main go routine] Finish importing posts")
	fmt.Println("[Main go routine] ---- waiting ----")
	<-done

	formatUserPosts(user, posts)
	fmt.Println("Done!!")
	fmt.Printf("User %v\n", user)
	for _, post := range posts {
		fmt.Printf("Post %v\n", post)
	}
}

func formatUserPosts(user *User, posts []*Post) {
	fmt.Println("[Main go routine] Start merging user posts")
	for _, post := range posts {
		post.UserId = user.Id
	}
	fmt.Println("[Main go routine] Finished merging user posts")
}

func importingPosts() []*Post {
	time.Sleep(1 * time.Second)
	titles := []string{"Post 1", "Random Post", "Second Post"}
	posts := []*Post{}
	for _, title := range titles {
		posts = append(posts, &Post{Title: title, Status: "draft"})
	}

	return posts
}

func buildingUser(user *User) {
	time.Sleep(2 * time.Second)
	user.Name = "Hello"
	user.LastName = "World"
	user.Status = "active"
	user.Id = "1"
}

func setDefaultTags(user *User) {
	time.Sleep(1 * time.Second)

	for _, tagName := range defaultags {
		user.Tags = append(user.Tags, &Tag{Name: tagName, Type: "System"})
	}
}
