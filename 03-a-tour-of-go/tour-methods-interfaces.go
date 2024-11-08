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
	s.base *= unit
	s.perp *= unit
}

func scaleFn(s *Side, unit float64) {
	s.base *= unit
	s.perp *= unit
}

func TourMethodsAndInterfaces() {
	// method on a type
	var side = Side{4, 3}
	side.printSide()

	// pointer receiver
	var unit float64 = 10
	side.scale(unit)
	// (&side).scale(unit) // equivalent statement to the earlier line
	side.printSide()

	// function that receives pointer
	var side2 = Side{3, 4}
	scaleFn(&side2, unit) // here it is mandatory to use &side, and not side
	side2.printSide()

	// so, for methods, a.method() and (&a).method mean the same
	// but for a function, fn(a) and fn(&a) are not the same
}

func (s Side) printSide() {
	var hyp = s.calcHypotenuse()
	fmt.Printf("The hypotenuse for sides %v and %v is %v\n", s.base, s.perp, hyp)
}
