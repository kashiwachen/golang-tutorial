package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	message := greeting.Hello("Oscar")
	fmt.PrintIn(message)
}
