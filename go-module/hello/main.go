package main

import (
	"fmt"
	"log"
	"vrouilhac/go-tutorials/go-module/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{
		"Vincent",
		"David",
		"Olivia",
		"Michael",
		"Julie",
	}

	msgs, err := greetings.Greetings(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range msgs {
		fmt.Println(msg)
	}
}
