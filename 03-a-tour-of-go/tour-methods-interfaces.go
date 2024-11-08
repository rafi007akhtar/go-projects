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

func TourMethodsAndInterfaces() {
	var side = Side{4, 3}
	var hyp = side.calcHypotenuse()
	fmt.Printf("The hypotenuse for sides %v and %v is %v\n", side.base, side.perp, hyp)
}
