package main

import (
	"fmt"
	"time"
)

func server1(ch chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		ch <- "from server1"
	}
}

func server2(ch chan<- string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "from server2"
	}
}

func main() {

	fmt.Println("Hello, playground")
	fmt.Println("-----------------")

	chanel1 := make(chan string)
	chanel2 := make(chan string)

	go server1(chanel1)
	go server2(chanel2)

	for {

		select {
		case s1 := <-chanel1:
			fmt.Println("Case 1, ", s1)
		case s2 := <-chanel2:
			fmt.Println("Case 2, ", s2)
		case s3 := <-chanel1:
			fmt.Println("Case 3, ", s3)
		case s4 := <-chanel2:
			fmt.Println("Case 4, ", s4)
		}

	}
}
