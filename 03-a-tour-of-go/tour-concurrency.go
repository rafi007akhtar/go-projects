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

func channelizeSlice(arr []int, ch chan int) {
	for _, val := range arr {
		ch <- val
	}
	close(ch)
}

func printChannelValues[T any](ch chan T) {
	println("Printing channel values:")
	for {
		var val, ok = <-ch
		if !ok {
			println("\nAll values of the given channel are printed")
			return
		}
		fmt.Printf("%v, ", val)
	}
}

func isEven(ch chan int, stop chan bool) {
	var val int
	for {
		select {
		// NOTE: I don't fully understand this
		case val = <-ch:
			if val%2 == 0 {
				fmt.Printf("%v is even\n", val)
			} else {
				fmt.Printf("%v is odd\n", val)
			}
		case <-stop:
			return
		default:
			println("Channels don't have value right now")
			return
		}
	}
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

	// 03 - buffered channels
	// to make a BUFFERED channel (with max capacity), apply the buffer length in the second argument of make
	var buffCh = make(chan int, 2)
	buffCh <- 1
	buffCh <- 2
	// buffCh <- 3 // this line will cause fatal error

	// 04 - closing a channel
	// if the channel is closed (ok value is false), no more value will be received from it
	var arr = []int{1, 2, 3, 4, 5}
	var myCh = make(chan int, 5)
	go channelizeSlice(arr, myCh)
	printChannelValues(myCh)

	// 05 - select statements
	var (
		ch2  = make(chan int)
		stop = make(chan bool)
	)
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		stop <- true
	}()
	// The following two function calls happen to the isEven function
	// one of them (could be either) will trigger the first case inside select until ch2 has value; then it will trigger the stop case
	// the other will trigger the default case as neither ch2 nor stop will have value then
	isEven(ch2, stop)
	isEven(ch2, stop)
}
