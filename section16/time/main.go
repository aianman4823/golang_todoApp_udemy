package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now()
	// fmt.Println(t)

	// t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	// fmt.Println(t2)
	// fmt.Println(t2.Year())
	// fmt.Println(t2.YearDay())
	// fmt.Println(t2.Month())
	// fmt.Println(t2.Weekday())
	// fmt.Println(t2.Day())
	// fmt.Println(t2.Hour())
	// fmt.Println(t2.Minute())
	// fmt.Println(t2.Second())
	// fmt.Println(t2.Nanosecond())
	// fmt.Println(t2.Zone())

	// time.Duration型
	// fmt.Println(time.Hour)
	// fmt.Printf("%T\n", time.Hour)
	// fmt.Println(time.Minute)
	// fmt.Println(time.Second)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Microsecond)
	// fmt.Println(time.Nanosecond)

	// //time.Duration型から文字列を生成
	// d, _ := time.ParseDuration("2h30m")
	// fmt.Println(d)

	// t3 := time.Now()
	// t3 = t3.Add(2*time.Minute + 15*time.Second)
	// fmt.Println(t3)

	// 時間の差分を取得
	// t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	// t6 := time.Now()
	// fmt.Println(t6)

	// // t5 - t6
	// d2 := t5.Sub(t6)
	// fmt.Println(d2)

	// // 時刻を比較
	// fmt.Println(t6.Before(t5))
	// fmt.Println(t6.After(t5))
	// fmt.Println(t5.Before(t6))
	// fmt.Println(t5.After(t6))

	// 指定時間のスリープ
	// 5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("５秒間停止")
}
