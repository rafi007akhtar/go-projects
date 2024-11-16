package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
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

// The following few functions are a solution to the equivalent trees exercise from the official Tour
// this one: https://go.dev/tour/concurrency/8
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var w func(t *tree.Tree)
	w = func(t *tree.Tree) {
		if t != nil {
			w(t.Right)
			ch <- t.Value
			w(t.Left)
		}
	}
	w(t)
}
func Same(t1, t2 *tree.Tree) bool {
	var (
		ch1 = make(chan int)
		ch2 = make(chan int)
	)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for val := range ch1 {
		if val != <-ch2 {
			return false
		}
	}
	return true
}
func printTree(t *tree.Tree) {
	if t != nil {
		printTree(t.Right)
		fmt.Printf("%v ", t.Value)
		printTree(t.Left)
	}
}
func testTreeEquivalence() {
	var (
		t1 = tree.New(1)
		t2 = tree.New(1) // change parameter anything else for inequivalence
	)
	println("First tree:")
	printTree(t1)
	println()
	println("Second tree:")
	printTree(t2)
	println()

	var same = Same(t1, t2)
	if same {
		println("Both trees are equivalent")
	} else {
		println("Both trees are not equivalent")
	}
}

// The following is a struct with a key/value entry, and a mutex to allow / disallow its modification
type SafeCounter struct {
	mut   sync.Mutex
	entry map[string]int
}

func (sc *SafeCounter) incr(key string) {
	sc.mut.Lock() // now only one goroutine can access it
	sc.entry[key]++
	sc.mut.Unlock() // now its free for use for other goroutines
}
func (sc *SafeCounter) get(key string) int {
	sc.mut.Lock()         // now only one goroutine can access it
	defer sc.mut.Unlock() // once the return is done, it will be free for use for other goroutines
	return sc.entry[key]
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
	// one of them (could be either) will trigger the first case inside select as long as ch2 has value; then it will trigger the stop case
	// the other will trigger the default case as neither ch2 nor stop will have value then
	isEven(ch2, stop)
	isEven(ch2, stop)

	// 06 - exercise: equivalent trees
	testTreeEquivalence()

	// 07 - go routines
	var counter = SafeCounter{entry: make(map[string]int)}
	const MY_KEY = "myKey"
	for i := 0; i < 1000; i++ {
		go counter.incr(MY_KEY)
	}
	// wait for a while until all lock / unlock operations are completed
	time.Sleep(time.Second)
	// now print the value
	fmt.Printf("Counter value at key '%v' is %v\n", MY_KEY, counter.get(MY_KEY))

	// 08 - exercise: web crawler
	fetchedUrls.urls = []string{}
	println("Web crawler results:")
	go Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
}

// the following code is a solution for the web crawler exercise on the official Tour
// this one: https://go.dev/tour/concurrency/10
type safeUrls struct {
	urls []string
	mut  sync.Mutex
}

var fetchedUrls safeUrls

func exists(searchTerm string, slice []string) bool {
	for _, val := range slice {
		if val == searchTerm {
			return true
		}
	}
	return false
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	fetchedUrls.mut.Lock()         // lock the url cache
	defer fetchedUrls.mut.Unlock() // unlock after function terminates

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !exists(u, fetchedUrls.urls) { // make sure duplicates are not refetched
			fetchedUrls.urls = append(fetchedUrls.urls, u)
			go Crawl(u, depth-1, fetcher) // parallelize here
		}
	}
	return
}

// <all the below code is all taken from the Go page>
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
