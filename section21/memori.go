package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)

	return s.Sys
}

func main() {
	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch
	}

	const numGorutines = 1000000

	wg.Add(numGorutines)

	before := memConsumed()

	for i := 0; i < numGorutines; i++ {
		go noop()
	}

	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGorutines/1000)
}
