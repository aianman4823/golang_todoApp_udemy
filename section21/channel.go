package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello"
	}()

	// waitgroupを使わずにチャネルの読み込みが遅延されている
	fmt.Println(<-ch)
}
