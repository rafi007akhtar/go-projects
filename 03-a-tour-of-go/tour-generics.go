package main

import "fmt"

// type parameters lie between function name and argement list, declared within []
// here, T is a type parameter that satisifes the contraint "comparable"
func indexOf[T comparable](slice []T, toFind T) int {
	for ind, val := range slice {
		if val == toFind {
			return ind
		}
	}
	return -1
}

// this linked list is a generic type
type List[T any] struct {
	val  T
	next *List[T]
}

// this function traverses a linked list
func traverseList[T any](list List[T]) {
	if list.next == nil {
		println("Linked list is empty")
		return
	}

	fmt.Printf("Value of node = %v\n", list.val)
	list = *list.next
	for {
		fmt.Printf("Next value of node = %v\n", list.val)
		if list.next == nil {
			break
		}
		list = *list.next
	}
}

func TourGenerics() {
	// 01 - type parameters
	var arr = []int{1, 2, 3, 4, 5}
	toFind := []int{2, 7}
	for _, val := range toFind {
		var ind = indexOf(arr, val)
		if ind != -1 {
			fmt.Printf("%v in the array exists at %v position", val, ind)
		} else {
			fmt.Printf("%v does not exist in the array", val)
		}
		println()
	}

	// 02 - generic types (structs with a parameter type of `any` constraint)
	var myList List[int]
	myList.val = 1
	myList.next = new(List[int])
	myList.next.val = 2
	myList.next.next = new(List[int])
	myList.next.next.val = 3
	myList.next.next.next = nil
	traverseList(myList)
}
