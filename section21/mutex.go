package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var wg sync.WaitGroup
	var data int

	// 順番とデータ競合が保障されていない
	// クリティカルセクション
	wg.Add(1)
	go func() {
		defer wg.Done()
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	wg.Wait()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(data)
	}
	memoryAccess.Unlock()
}
