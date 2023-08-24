package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()

}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"deta",
		"epsilon",
	}

	wg.Add(13)
	for i, x := range words {

		go printSomething(fmt.Sprintf("%d:%s", i, x), &wg)

	}
	wg.Wait()

	//fmt.Println("hello world")
	//go printSomething("this  is start")
	//time.Sleep(1 * time.Second) //badway
	//wg.Wait()
	fmt.Println("this is second time we call")
}
