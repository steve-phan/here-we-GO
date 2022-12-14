package main

import (
	"log"
	"os"

	"bookable24.de/main/handlejson"
	"bookable24.de/main/helpers"
	"github.com/joho/godotenv"
)

// "fmt"
// "log"

// "bookable24.de/main/greetings"
// "bookable24.de/main/overviewgo"

const numPool = 500

func CaculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// log.SetPrefix("Show:   ")
	log.SetFlags(0)

	// A slice of names.
	// names := []string{"Teddy", "Luna", "Lily"}

	// Request greeting messages for the names slice
	// messages, err := greetings.Hellos(names)

	// Request a greeting message.
	// message, err := greetings.Hello("Luna")
	// If an error was returned, print it to the console and
	// exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If no error was returned, print the returned
	// message | messages to the console.
	// fmt.Println(message)
	// fmt.Println(messages)

	// call overviewgo module
	// overviewgo.ShowMyMap()
	// overviewgo.ShowMySlice()
	// overviewgo.ShowInfo()

	// intChan := make(chan int)
	// defer close(intChan)

	// go CaculateValue(intChan)

	// log.Println(<-intChan)
	handlejson.ParseJson()

	// Loading env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error, loading env fail")
	}
	// Get Key
	key := os.Getenv("SUPER_KEY")
	log.Println("KEY", key)

}
