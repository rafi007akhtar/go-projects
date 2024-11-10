package main

import (
	"fmt"
	"io"
	"math"

	"golang.org/x/tour/reader"
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

// Implementing Stringer interface on Side
func (s Side) String() string {
	var hyp = s.calcHypotenuse()
	return fmt.Sprintf("The hypotenuse for sides %v and %v is %v\n", s.base, s.perp, hyp)
}

// NOTE: the following type, method and function are my solution to the Errors exercise on Tour
// https://go.dev/tour/methods/20
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	var floatE = float64(e) // NOTE: this is absolutely necessary, to convert e to float value before referencing in a print function
	return fmt.Sprintf("cannot Sqrt negative number: %v", floatE)
}
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeSqrt(n)
	}
	var z = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - n) / (2 * z)
	}

	var prevZ float64
	for {
		prevZ = z
		z -= (z*z - n) / (2 * z)
		if (z - prevZ) < 0.000001 {
			break
		}
	}
	return z, nil
}

// NOTE: this is my solution to the Readers exercise on the official Tour page
// this one: https://go.dev/tour/methods/22
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

func (r MyReader) Read(data []byte) (int, error) {
	for i := 0; i < len(data); i++ {
		data[i] = 'A'
	}
	return len(data), nil
}
func testReader() {
	reader.Validate(MyReader{})
}

// NOTE: the following is a solution to the rot13 exercise on the official Tour page
// https://go.dev/tour/methods/23
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	var n, err = r.r.Read(b)
	for i, c := range b {
		if (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z') {
			b[i] -= 13
		} else if (c >= 'A' && c <= 'M') || (c >= 'a' || c <= 'm') {
			b[i] += 13
		}
	}
	return n, err
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

	// 05 - type assertions and switch
	// we can provide type assertion of a variable this way:
	var a, ok = any.(Side)
	if ok {
		fmt.Printf("any variable is of type Side with value %v\n", a)
	}
	// NOTE: the second variable in the assert (ok) is optional, but it's better to provide it otherwise the compiler will panic on assert fail
	aa, ok := any.(int)
	if !ok {
		fmt.Printf("any variable is not of type int, so we get its zero-value %v\n", aa)
	}
	// you can put type in a switch case with variable.(type) syntax, used below:
	var some, random, variables interface{}
	some = 1
	random = "hello"
	variables = Side{}
	doStuffFromType([]interface{}{some, random, variables})

	// 06 - Stringer interface
	// Stringer interface is implemented on the Side type
	var side5 = Side{7, 8}
	fmt.Print(side5)

	// 07 - errors
	var n float64 = 2
	var (
		sqrt1, err1 = Sqrt(n)
		sqrt2, err2 = Sqrt(-n)
	)
	printSqrt(n, sqrt1, err1)
	printSqrt(n, sqrt2, err2)

	// 08 - readers
	fmt.Print("Reader solution is: ")
	testReader()
}

func (s Side) printSide() {
	var hyp = s.calcHypotenuse()
	fmt.Printf("The hypotenuse for sides %v and %v is %v\n", s.base, s.perp, hyp)
}

func printAny(any interface{}) {
	fmt.Printf("Variable of type %T has value %v\n", any, any)
}

func doStuffFromType(someVars []interface{}) {
	for _, someVar := range someVars {
		switch v := someVar.(type) {
		case int:
			fmt.Printf("Double of %v is %v\n", v, v*2)
		case string:
			fmt.Printf("Variable %v is a string\n", v)
		default:
			fmt.Printf("Variable is of type %T\n", v)
		}
	}
}

func printSqrt(n float64, root float64, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v ** 0.5 = %v\n", n, root)
	}
}
