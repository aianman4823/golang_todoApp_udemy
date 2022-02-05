package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup
	// m := map[string]int{"A": 0, "B": 1}

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		defer wg.Done()
	// 		m["A"] = rand.Intn(100)
	// 		m["B"] = rand.Intn(100)
	// 	}()

	// 	go func() {
	// 		defer wg.Done()
	// 		m["A"] = rand.Intn(100)
	// 		m["B"] = rand.Intn(100)
	// 	}()
	// }

	// syncのmap型の定義
	smap := &sync.Map{}

	// syncのmap型の使い方
	smap.Store("Hello", "World")
	smap.Store(1, 2)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	smap.Delete(1)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	v, ok := smap.Load("Hello")
	if ok {
		fmt.Println(v)
	}

	smap.LoadOrStore("Hello", "Wooooorld")
	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	smap.LoadOrStore(2, 3)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
