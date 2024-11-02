package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	var msg = greetings.Hello("Rafi")
	fmt.Println(msg)
}
