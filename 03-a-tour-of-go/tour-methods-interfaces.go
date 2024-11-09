package main

import (
	"fmt"
	"math"
)

type Side struct{ base, perp float64 }

// Here, we are defining a method `calcHypotenuse` on the type `Sides`
func (s Side) calcHypotenuse() float64 {
	return math.Sqrt(s.base*s.base + s.perp*s.perp)
}

// func (s Side) Scale(unit float64) { // this passes by value, and works on copied values of the object, thus not changing it
func (s *Side) scale(unit float64) { // this passes by address, and changes the values of the calling object
	if s == nil {
		fmt.Printf("Cannot scale a nil side\n")
		return
	}
	s.base *= unit
	s.perp *= unit
}

func scaleFn(s *Side, unit float64) {
	s.base *= unit
	s.perp *= unit
}

type rightTri interface {
	calcHypotenuse() float64
	printSide()
}

type scalableRightTri interface {
	scale(unit float64)
}

func TourMethodsAndInterfaces() {
	// 01 - method on a type
	var side = Side{4, 3}
	side.printSide()

	// 02 - pointer receiver
	var unit float64 = 10
	side.scale(unit)
	// (&side).scale(unit) // equivalent statement to the earlier line
	side.printSide()

	// 03 - function that receives pointer
	var side2 = Side{3, 4}
	scaleFn(&side2, unit) // here it is mandatory to use &side, and not side
	side2.printSide()
	// so, for methods, a.method() and (&a).method mean the same
	// but for a function, fn(a) and fn(&a) are not the same

	// 04 - interfaces
	// here, a method gets mentioned in the interface
	// a variable implements an interface, by using the interface as a type
	// the variable is then assigned a value
	// the type of the value MUST implement the method(s) mentioned in the interface
	var side3 rightTri
	side3 = Side{1, 3} // this works because type of side3 (Side) implements the calcHypotenuse method
	side3.printSide()

	// this also extends to methods that receive pointer arguments
	var side4 scalableRightTri // side4 implements such an interface
	side4 = &Side{1, 3}        // so here it must be &Side, and not Side directly
	fmt.Printf("side4 before scaling = %v;\t", side4)
	side4.scale(5)
	fmt.Printf("side4 after scaling = %v\n", side4)
	var nilSide *Side
	side4 = nilSide
	side4.scale(5) // nil will be handled gracefully here

	// an empty interface (interface{}) is Go's equivalent to Typescript's any type
	var any interface{}
	any = 1
	printAny(any)
	any = "hello"
	printAny(any)
	any = Side{3, 4}
	printAny(any)
}

func (s Side) printSide() {
	var hyp = s.calcHypotenuse()
	fmt.Printf("The hypotenuse for sides %v and %v is %v\n", s.base, s.perp, hyp)
}

func printAny(any interface{}) {
	fmt.Printf("Variable of type %T has value %v\n", any, any)
}
