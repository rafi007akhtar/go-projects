package main

import "fmt"

func main() {
	var choice int

	println("Enter: \n1 for basics, \n2 for methods and interpreters \nYour choice:")
	fmt.Scanf("%d\n", &choice)

	switch choice {
	case 1:
		TourBasics()
	case 2:
		TourMethodsAndInterfaces()
	default:
		println("This choice is not valid")
	}
}
