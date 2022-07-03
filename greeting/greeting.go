// Declare greeting package, and hold the related functions
package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
