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

func randomFormat() string {
	greetings := []string{
		"Olá, %v!",
		"Qualé, %v! Tudo em cima?",
		"Como é que é, %v? Numa  boa?",
	}

	return greetings[rand.Intn(len(greetings))]
}
