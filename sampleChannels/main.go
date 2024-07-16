package main

import (
	"fmt"
	"strings"
)

func main() {
	//create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	for {

		fmt.Println("Enter a message ou type Q to quit and ENTER:")
		var msg string
		fmt.Scanln(&msg)

		if strings.ToUpper(msg) == "Q" {
			break

		}
		ping <- msg
		//wait for response
		response := <-pong
		fmt.Println(response)
	}

	fmt.Println("Bye, World!")
}

func shout(ping <-chan string, pong chan<- string) {

	for {
		msg := <-ping
		pong <- strings.ToUpper(fmt.Sprintf("%s !!!", msg))
	}
}
