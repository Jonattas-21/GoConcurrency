package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	printsomething("Test", &waitGroup)
	waitGroup.Wait()

	w.Close()
	output, _ := io.ReadAll(r)
	result := string(output)

	os.Stdout = stdOut

	if !strings.Contains(result, "Test") {
		t.Errorf("Expected 'Test', got '%s'", result)
	}
}
