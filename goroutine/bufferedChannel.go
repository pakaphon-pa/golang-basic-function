package goroutine

import (
	"fmt"
	"time"
)

func BufferedChannel() {
	names := make(chan string, 3)
	go generateName(names)

	time.Sleep(5 * time.Second)

	for name := range names {
		fmt.Printf("Name received: %v\n", name)
	}
}

func generateName(names chan<- string) {
	defer close(names)

	for _, name := range []string{"Carl", "Paul", "May", "Laura", "John"} {
		time.Sleep(1 * time.Second)
		names <- name
		fmt.Printf("Name sent: %v\n", name)
	}
}
