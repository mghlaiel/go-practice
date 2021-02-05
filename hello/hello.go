package main

import (
	"fmt"
	"log"

	"github.com/mghlaiel/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"nero", "valerie", "mike"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
