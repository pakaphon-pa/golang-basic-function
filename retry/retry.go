package retry

import (
	"log"
	"time"
)

var functionPassing bool = false
var attempts int = 5

func retry(attemps int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attemps; i++ {
		if i > 0 {
			log.Println("retrying after error: ", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f()
		if err == nil {
			return nil
		}
	}

	return err
}
