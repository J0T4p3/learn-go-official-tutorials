package main

import (
	"log"
	"fmt"
	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	message, err:= greetings.Hello("João")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
