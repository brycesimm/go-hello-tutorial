package main

// Run go run . to run main function
// go build will create a .exe in this directory
// go install will crate a .exe in the C:\Users\{username}\go\bin directory where it can be called globally

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Chip")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	names := []string{"Chip", "Dave", "Tanya"}

	messages, err2 := greetings.Hellos(names)

	if err2 != nil {
		log.Fatal((err2))
	}

	fmt.Println(messages)
}
