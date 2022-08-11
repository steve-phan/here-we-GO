package main

import (
	"fmt"
	"time"
)

// Print the n number of fibonacci number

func main() {
	number := 100
	numberOfWorker := 8
	jobs := make(chan int, number)
	results := make(chan int, number)

	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}

	now := time.Now().UnixMilli()
	fmt.Printf("Now is %v\n", now)

	for i := 0; i < number; i++ {
		jobs <- i
		// fmt.Println(<-results)
		// fmt.Println(fibonacci(i))
		// fmt.Println("Hello")
		if <-results == 4807526976 {
			end := time.Now().UnixMilli()
			fmt.Printf("Now is %v\n", end)
			fmt.Printf("Time running is %v", end-now)
		}
	}
	// 38926
	// 37484 37322
	// end := time.Now().UnixMilli()
	// fmt.Printf("Now is %v\n", end)
	// fmt.Printf("Time running is %v", end-now)
	close(jobs)

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}

}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
