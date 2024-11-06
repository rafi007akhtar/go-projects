package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"golang.org/x/tour/pic"
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

type Vertex struct {
	x int
	y int
}

func sliceProps() {
	var primes = []int{2, 3, 5, 7, 11, 13}
	var lenp, capp = len(primes), cap(primes)
	fmt.Printf("For slice %v, length is %v and capacity is %v\n", primes, lenp, capp)

	primes = primes[:0]
	lenp, capp = len(primes), cap(primes)
	fmt.Printf("Shrinking the primes slice so it now becomes %v with %v length and %v capacity\n", primes, lenp, capp)
	fmt.Printf("It is %t that this slice is nil\n", primes == nil)

	primes = primes[:4] // NOTE: primes[n:] would not work if len(primes) < n
	lenp, capp = len(primes), cap(primes)
	fmt.Printf(
		"Now that the slice is empty, re-expanding it by slicing on its original array so it it now becomes %v with length %v and capacity %v\n",
		primes,
		lenp,
		capp,
	)

	var undef []int
	fmt.Printf("Created an undefined slice, %v, so it's %t that the zero-value of this slice is nil\n", undef, undef == nil)

	var (
		a = make([]int, 5)
		b = make([]int, 0, 5)
		c = b[:2]
		d = c[2:5]
	)
	var slices = [][]int{a, b, c, d}
	var e = []int{1, 2, 3, 4, 5}
	slices = append(slices, e)
	printSlices(slices)
}

func printSlice(s []int) {
	fmt.Printf("For slice %v, length is %v and capacity is %v\n", s, len(s), cap(s))
}

func printSlices(slices [][]int) {
	for i := 0; i < len(slices); i++ {
		printSlice(slices[i])
	}
}

func arrs() {
	var arr1 = [9]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy"}
	fmt.Println("arr1:", arr1, arr1[len(arr1)-1])
	arr1[8] = "dog"
	fmt.Println("arr1:", arr1)
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

	// 11 - struct literals
	var v1 = Vertex{1, 2}
	var v2 = Vertex{x: 1}
	var v3 = Vertex{}
	var vp = &Vertex{1, 2}
	fmt.Printf("v1 = %+v, v2 = %+v, v3 = %+v, vp = %+v, vp value = %+v\n", v1, v2, v3, vp, *vp)

	// 12 - arrays
	arrs()
	println()

	// 13 - slices and ranges
	sliceProps()
	pic.Show(Pic)
}
