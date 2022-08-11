package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2) // assign a number(N) of wg tasks

	go func() {
		countLeng("Hello World")
		wg.Done() // it means N--
	}()

	go func() {
		countLeng("GoodBye")
		wg.Done() // it means N-- = 0
	}()

	wg.Wait() // When N == 0 (finish and go next)
	fmt.Println("Count done")

}

func countLeng(str string) {

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(string(str[i]))
	// }
	for _, char := range str {
		fmt.Println(string(char))
	}

}
