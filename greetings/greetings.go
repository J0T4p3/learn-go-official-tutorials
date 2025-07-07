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
	// Verifies if the map is empty
	if len(names) == 0 {
		return nil, errors.New("empty name")
	}

	// Declares the new map (dictionary) of keys and values as string
	messages := make(map[string]string)

	// Iterates over the map of names and create a new message for
	// each of them
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	greetings := []string{
		"Olá, %v!",
		"Qualé, %v! Tudo em cima?",
		"Como é que é, %v? Numa  boa?",
	}
	return greetings[rand.Intn(len(greetings))]
}
