package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("When was you born?: ")
	scanner.Scan()
	input := scanner.Text()

	inputNum, _ := strconv.ParseInt(input, 10, 64)

	log.Printf("%d is the input", 2022-inputNum)

}
