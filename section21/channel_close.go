package main

import (
	"fmt"
	"sync"
)

func main() {
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello"
	// }()

	// v, ok := <-ch
	// fmt.Println(v, ok)
	// close(ch)
	// v, ok = <-ch
	// fmt.Println(v, ok)

	// ch := make(chan int)

	// go func() {
	// 	defer close(ch)

	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// }()

	// for integer := range ch {
	// 	fmt.Println(integer)
	// }

	begin := make(chan interface{})

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		fmt.Printf("Start goroutine %d\n", i)

		go func(i int) {
			defer wg.Done()
			<-begin

			fmt.Printf("%d has begin\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutine!")

	// ちゃねるのcloseを利用するとgo routineにシグナルを送ることができる
	close(begin)

	wg.Wait()
}
