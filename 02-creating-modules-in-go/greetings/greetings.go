package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// allocate memory for a map with string keys and string values
	var messages = make(map[string]string)

	// loop through the range of names; ignore the index with an underscore
	for _, name := range names {
		// get the greeting, set the message for the name (with error checking)

		var message, err = Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	// return the messages map
	return messages, nil
}

func randomFormat() string {
	var formats = []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Greetings, %v. We finally meet.",
	}

	return formats[rand.Intn(len(formats))]
}
