package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple execution
	// Output will be in the order
	example1("Hello world 1")
	example1("Hello world 2")

	// Second execution
	// Here you will not see the output of Hello world 3 because main goroutine has already finished
	// and not wait for the 'result' of this goroutine
	go example1("Hello world 3")
	example1("Hello world 4")

	// Third execution
	// Here you can see the output because sleep time allows to wait the main goroutine
	// and the execution of this goroutine
	// This is not the way to do this, will be demonstrate of the following step
	go example1("Hello world 5")
	time.Sleep(1 * time.Second)
	example1("Hello world 6")
}

func example1(s string) {
	fmt.Println(s)
}
