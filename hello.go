// Declare the main package.
package main

// Import built-in package
import (
	"fmt"

	"rsc.io/quote"
)

// Implement a main function, which is executed by default when main package is run
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
