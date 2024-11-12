package main

import (
	"fmt"
	"time"
)

func say(word string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		println(word)
	}
}

/*
Input a slice of ints and a channel
Sum the slice and send to the channel
*/
func sumCh(nums []int, ch chan int) {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum // NOTE: this is where we send the value to the channel
}

func TourConcurrency() {
	// 01 - goroutines
	go say("hello", 5)
	say("world", 5)

	// 02 - channels
	// from GFG website: In Go language,
	// a channel is a medium through which a goroutine communicates with another goroutine
	// and this communication is lock-free.
	// Task: split an slice in two, sum each half in a channel, print both the values and their sum
	var nums = []int{7, 2, 8, -9, 4, 0}
	var firstHalf = nums[:len(nums)/2]
	var secondHalf = nums[len(nums)/2:]
	// now allocate a channel, and create two goroutines where the channel will be involved
	var ch = make(chan int)
	go sumCh(firstHalf, ch)
	go sumCh(secondHalf, ch)
	// now extract the values put into these channels into variables
	var (
		x = <-ch
		y = <-ch
	)
	fmt.Printf("Extracted values: %v, %v; their sum: %v\n", x, y, x+y)
}
