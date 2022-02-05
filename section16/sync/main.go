package main

import (
	"fmt"
	"sync"
	"time"
)

var st struct {
	A, B, C int
}

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	// アンロック
	mutex.Unlock()
}

// ミューテックスによる同期処理
func main() {
	// mutex = new(sync.Mutex)

	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		for i := 0; i < 1000; i++ {
	// 			UpdateAndPrint(i)
	// 		}
	// 	}()
	// }

	// for {

	// }

	// ゴルーチン終了を待ち受ける
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2st Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3st Goroutine")
		}
		wg.Done()
	}()
	wg.Wait()
}
