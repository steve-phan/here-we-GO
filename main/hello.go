package main

import (
	"fmt"
	"log"

	"bookable24.de/greetings" // TODO : $GOPATH => BrokenImport
	"bookable24.de/overviewgo"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Show:   ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Teddy", "Luna", "Lily"}

	// Request greeting messages for the names slice
	messages, err := greetings.Hellos(names)

	// Request a greeting message.
	// message, err := greetings.Hello("Luna")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned
	// message | messages to the console.
	// fmt.Println(message)
	fmt.Println(messages)

	// call overviewgo module
	overviewgo.ShowMyMap()
	overviewgo.ShowMySlice()

}
