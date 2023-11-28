package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething2(t *testing.T) {

	// Way to capture os.Stdout
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execution of the function to test
	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething2("test", &wg)
	wg.Wait()

	// back to normal state
	_ = w.Close()
	os.Stdout = old // restoring the real stdout

	outputBytes, _ := io.ReadAll(r)
	output := string(outputBytes)

	if !strings.Contains(output, "test") {
		t.Errorf("Output should be test")
	}

}
