package main

import (
	"fmt"
	"sync"
)

func main() {
	words := []string{
		"Alpha",
		"Beta",
		"Gamma",
		"Delta",
		"Epsilon",
		"Zeta",
		"Eta",
		"Theta",
	}

	var waitGoup sync.WaitGroup

	waitGoup.Add(len(words))

	for i, item := range words {
		go printsomething(fmt.Sprintf("%d = %s", i, item), &waitGoup)
	}

	waitGoup.Wait()
	waitGoup.Add(1)
	printsomething("Done, all items printed", &waitGoup)
}

func printsomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
