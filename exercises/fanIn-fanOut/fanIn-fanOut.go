package main

import "fmt"

const sqrtNumber = 4

func main() {

	numbers := []int{}

	for i := 0; i < sqrtNumber; i++ {
		numbers = append(numbers, i)
	}

	// Generate PipeLine
	inputChan := generatePipeLine((numbers))

	// FanOut
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)
	c5 := fanOut(inputChan)
	c6 := fanOut(inputChan)
	c7 := fanOut(inputChan)
	c8 := fanOut(inputChan)
	// FanIn

	c := fanIn(c1, c2, c3, c4, c5, c6, c7, c8)

	sum := 0
	for i := 0; i < 4; i++ {
		sum = sum + <-c
	}

	fmt.Printf("Them sqrt number of %v is %v", sqrtNumber, sum)
}

func generatePipeLine(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChan ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChan {
			for n := range c {
				fmt.Println(n)
				in <- n
			}
		}
	}()
	return in
}
