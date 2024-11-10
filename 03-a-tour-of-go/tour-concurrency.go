package main

import "time"

func say(word string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		println(word)
	}
}

func TourConcurrency() {
	// 01 - goroutines
	go say("hello", 5)
	say("world", 5)
}
