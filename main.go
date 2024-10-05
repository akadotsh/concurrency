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

	// Adding WaitGroup for fibonacci goroutine
	wg.Add(1)
	go func() {
		fibonacci()
		defer wg.Done() // Calling Done when fibonacci func is finished executing
	}()

	
	wg.Wait()


}

func Counter(message chan string) {
	time.Sleep(time.Second * 2)

	// using channels to send value
	message <- "Counter"
}

func fibonacci() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Printf("Jobs in channel: %d\n", len(jobs))

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	fmt.Println("CALLED")
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
