package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	var msg, err = greetings.Hello("Rafi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
