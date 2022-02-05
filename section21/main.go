package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func closure() {
	var wg sync.WaitGroup
	say := "Hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		say = "Good Bye"
	}()

	wg.Wait()

	fmt.Println(say)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello2")
	}()

	wg.Wait()
}
