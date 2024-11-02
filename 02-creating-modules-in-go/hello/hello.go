package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// logger settings
	log.SetPrefix("greetings: ") // set prefix
	log.SetFlags(0)              // don't log timestamp and source file info

	var names = []string{"Md", "Rafi", "Akhtar"}
	var messages, err = greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
