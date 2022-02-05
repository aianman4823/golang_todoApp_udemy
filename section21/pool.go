package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
}

func main() {
	count := 0
	mypool := &sync.Pool{
		New: func() interface{} {
			count++
			fmt.Println("Creating...")
			return struct{}{}
		},
	}

	mypool.Put("manualy added: 1")
	mypool.Put("manualy added: 2")

	var wg sync.WaitGroup

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Microsecond)
		go func() {
			defer wg.Done()
			instance := mypool.Get()
			mypool.Put(instance)
		}()
	}

	wg.Wait()

	fmt.Printf("created instance: %d", count)
}
