package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	tasks := []string{"A", "B", "C"}

	for _, task := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			fmt.Println(task)
		}(task)
	}
	wg.Wait()
}
