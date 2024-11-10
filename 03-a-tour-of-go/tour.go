package main

import "fmt"

func main() {
	var choice int

	println("Enter: \n1 for basics, \n2 for methods and interpreters \n3 for generics \n4 for goroutines \nYour choice:")
	fmt.Scanf("%d\n", &choice)
	if choice < 1 || choice > 4 {
		choice = 4
		println("We are going with the default choice, which is", choice)
	}

	switch choice {
	case 1:
		TourBasics()
	case 2:
		TourMethodsAndInterfaces()
	case 3:
		TourGenerics()
	case 4:
		TourConcurrency()
	default:
		println("This choice is not valid")
	}
}
