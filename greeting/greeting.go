// Declare greeting package, and hold the related functions
package greeting

import (
	"errors"
	"fmt"
)

// Implement Hello function
// Hello returns a greeting for the named person
// Declare the input and output
func Hello(name string) (string, error) {
	// Declare a message variable.
	// ":=" declare and init a variable
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
