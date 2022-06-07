package goroutine

import (
	"fmt"
	"time"
)

func GoRoutine() {
	fmt.Println("Main function")
	go countNumbers(20)
	time.Sleep(1 * time.Second)
	fmt.Println("End Main function")
}

func countNumbers(limit int) {
	num := 0
	for i := 1; i < limit; i++ {
		num += i
	}

	fmt.Println("Num: ", num)
}
