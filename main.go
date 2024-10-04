package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	message := make(chan string)

	wg.Add(1)
	go func() {
		Counter(message)
		defer wg.Done()
	}()

	fmt.Println("End of program " + <-message)

	wg.Wait()
}

func Counter(message chan string) {
	time.Sleep(time.Second * 2)

	// using channels to send value
	message <- "Counter"
}
