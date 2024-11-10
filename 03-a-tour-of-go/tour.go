package main

import "fmt"

func main() {
	var choice int

	println("Enter: \n1 for basics, \n2 for methods and interpreters \n3 for generics \nYour choice:")
	fmt.Scanf("%d\n", &choice)
	if choice != 1 && choice != 2 {
		choice = 3
		println("We are going with the default choice, which is", choice)
	}

	switch choice {
	case 1:
		TourBasics()
	case 2:
		TourMethodsAndInterfaces()
	case 3:
		TourGenerics()
	default:
		println("This choice is not valid")
	}
}
