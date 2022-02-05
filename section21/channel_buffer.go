package main

import (
	"fmt"
	"time"
)

func main() {
	// 5とbufferを利用しないと5回書き込みをしたのち、
	// channelをcloseするまでに読み込み側で時間がかかる場合は、書き込み側も遅延される
	ch := make(chan int, 5)

	go func() {
		defer fmt.Println("Close")
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Printf("writing to channel: %v\n", i)
			ch <- i
		}
	}()

	for integer := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("reading to channel: %v\n", integer)
	}
}
