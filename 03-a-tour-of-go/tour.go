package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func add(a, b int) int {
	return a + b
}

func twoRandomNumbers(n int) (r1, r2 int) {
	r1 = rand.Intn(n)
	r2 = rand.Intn(n)
	return
}

func factorial(n int) int {
	var fact = 1
	for i := 1; i < n; i++ {
		fact = fact * i
	}
	return fact
}

func fibonacci(n int) int {
	var a, b, c = 1, 1, 0
	var count = 2
	for count < n {
		c = a + b
		a, b = b, c
		count += 1
	}
	return c
}

func sqrt(n float64) float64 {
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
	return z
}

func getOS() string {
	switch os := runtime.GOOS; os {
	case "android":
		return "android"
	case "windows":
		return "windows"
	default:
		return os
	}
}

func helloWorld() {
	defer fmt.Println("world")
	println("hello")
}

func pointers() {
	var i, j = 10, 20
	var p1, p2 *int
	p1 = &i // p1 is assigned the address of i
	p2 = &j // p2 is assigned the address of j
	fmt.Printf("Address of i (%v) is %p\n", i, p1)
	fmt.Printf("Address of j (%v) is %p\n", j, p2)

	*p1 /= 2
	*p2 *= 2
	fmt.Printf("The updated values of i and j are %v and %v respectively", i, j)
}

type Coord struct {
	lat float64
	lng float64
}

func coordinates(x, y float64) Coord {
	var c = Coord{x, y}
	return c
}

func main() {
	// 01 - sum of two numbers, and print
	var a, b = 12, 12
	var sum = add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, sum)

	// 02 - "naked returns" - specified while declaring the function
	var r1, r2 = twoRandomNumbers(10)
	fmt.Printf("Two random numbers are %v and %v\n", r1, r2)

	// 03 - "zero value" - 0 for numeric types, "" for strings, false for bool
	var undef int
	fmt.Printf("The variable undef isn't defined explicitly but its value is %v\n", undef)

	// 04 - for loops
	var n = 7
	fmt.Printf("The factorial of %v is %v\n", n, factorial(n))

	// 05 - for as while loop
	n = 8
	fmt.Printf("The Fibonacci number %v is %v\n", n, fibonacci(n))

	// 06 - square root
	var x = 10.0
	fmt.Printf("Square root of %v is nearly equal to: %v\n", x, sqrt(x))

	// 07 - switch case
	fmt.Printf("The operating system is %v\n", getOS())

	// 08 - defer keyword
	// the code following defer won't be executed until the surrounding function returns
	// multiple defer statements are stacked and executed in LIFO order
	helloWorld()

	// 09 - pointers
	// & is "address of"
	// * is "value at"
	pointers()

	// 10 - structs
	var y float64
	x, y = 110, 120
	var c = coordinates(x, y)
	fmt.Printf("\nThe coordinates are %+v, where the latitude is %v and the longitude is %v\n", c, c.lat, c.lng)
}
