package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_update_Message(t *testing.T) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go updateMessagbe("Hello universe", &waitGroup)
	waitGroup.Wait()

	if msg != "Hello universe" {
		t.Errorf("Expected 'Hello universe', got '%s'", msg)
	}

}

func Test_main_print(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello World"
	printMessage()

	w.Close()
	output, _ := io.ReadAll(r)
	result := string(output)

	os.Stdout = stdOut

	if !strings.Contains(result, "Hello World") {
		t.Errorf("Expected 'world', got '%s'", result)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	output, _ := io.ReadAll(r)
	result := string(output)
	os.Stdout = stdOut

	if !strings.Contains(result, "Hello World") {
		t.Errorf("Expected 'Hello World', got '%s'", result)
	}

	if !strings.Contains(result, "Hello cosmos") {
		t.Errorf("Expected 'Hello cosmos', got '%s'", result)
	}

	if !strings.Contains(result, "Hello universe") {
		t.Errorf("Expected 'Hello universe', got '%s'", result)
	}
}
