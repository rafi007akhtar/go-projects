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

	var msg, err = greetings.Hello("Rafi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
