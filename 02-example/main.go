package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func printSomething2(s string, wg *sync.WaitGroup) {
	// When the function exits, decrease the waitGroup by one with wg.Done()
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// Second way with wait group
	way2()

	// Dummy way like first-example
	//way1()
}

// Second way with sync.waitGroup
func way2() {
	var wg sync.WaitGroup

	players := []string{
		"LeBron James",
		"Kobe Bryant",
		"Michael Jordan",
		"Luka Doncic",
		"Tony Parker",
		"Magic Johnson",
		"Tim Duncan",
	}

	for i, p := range players {
		// Add 1 goroutine in the waitGroup for the execution
		wg.Add(1)
		go printSomething2(fmt.Sprintf("Player %d is %s", i, p), &wg)
	}
	// Use wg.Wait to wait for all goroutine to finish
	wg.Wait()

	// Should place wg.Add(1) here because of wg.Done() is played in the function
	// Otherwise -> panic: sync: negative WaitGroup counter
	wg.Add(1)
	printSomething2("Test last execution", &wg)
}

// First way with time.Sleep which is very bad (we cannot know the time to wait for all goroutines)
func way1() {
	words := []string{
		"Alpha",
		"Beta",
		"Gamma",
		"Delta",
		"Epsilon",
		"Zêta",
		"Êta",
		"Thêta",
		"Iota",
		"Kappa",
	}

	// Very interesting, because this is not executed in order
	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word))
	}
	// Use sleep time to wait for goroutine
	time.Sleep(1 * time.Second)
}
