package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Oscar", "Helen", "Stella"}
	messages, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err) // return the error and exit the program
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
