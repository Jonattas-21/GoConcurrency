package main

import (
	"fmt"
	"sync"
)

var msg string

func printMessage() {
	fmt.Println(msg)
}

func updateMessagbe(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	msg = s
}

func main() {

	var waitGroup sync.WaitGroup

	msg = "Hello World"

	waitGroup.Add(1)
	go updateMessagbe("Hello universe", &waitGroup)
	printMessage()
	waitGroup.Wait()

	waitGroup.Add(1)
	go updateMessagbe("Hello cosmos", &waitGroup)
	printMessage()
	waitGroup.Wait()

	waitGroup.Add(1)
	go updateMessagbe("Hello World", &waitGroup)
	printMessage()
	waitGroup.Wait()
}
