// Declare greeting package, and hold the related functions
package main

import "fmt"

// Implement Hello function
// Hello returns a greeting for the named person
func Hello(name string) string {
	// Declare a message variable.
	// ":=" declare and init a variable
	message := fmt.Spriintf("Hi, %v. Welcome!", name)
	return message
}
