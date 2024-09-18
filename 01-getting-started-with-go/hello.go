package main

import (
	"fmt"

	"rsc.io/quote"
)

func phrase() {
	fmt.Println(quote.Go())
}

func main() {
	fmt.Println("Hello, world")
	phrase()
}
