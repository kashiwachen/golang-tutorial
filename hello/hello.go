package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	message, err := greeting.Hello("")

	if err != nil {
		log.Fatal(err) // return the error and exit the program
	}

	fmt.Println(message)
}
