package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int
	var i int = 100
	fmt.Println(i)
	var i2 int64 = 200
	// fmt.Println(i + i2)
	// Printfと%Tで型を調べられる
	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))

	// float
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// float64になる(intとは違って環境依存ではないので計算可能)
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// bool
	var t, f bool = true, false
	fmt.Println(t, f)

	// string
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
	test
	`)

	fmt.Println("\"")
	fmt.Println(`"`)

	// 文字列はバイト型の集まり
	fmt.Println(string(s[0]))

	// byte
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3)

	arr4 := [...]string{"C"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 中身の要素数を調べる
	fmt.Println(len(arr1))

	// interface & nil
	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = 3.4
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
	// x = 2
	// fmt.Println(x + 3)

	// 型変換
	var in int = 1
	f64 := float64(in)
	fmt.Println(f64)
	fmt.Printf("in = %T\n", in)
	fmt.Printf("f64 = %T\n", f64)

	in2 := int(f64)
	fmt.Printf("in2 = %T\n", in2)

	f2 := 10.5
	in3 := int(f2)
	fmt.Printf("in3 = %T\n", in3)
	fmt.Println(in3)

	var st string = ""
	fmt.Printf("s = %T\n", st)

	in3, _ = strconv.Atoi(st)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(in3)
	fmt.Printf("i = %T\n", in3)

	var in5 int = 200
	s6 := strconv.Itoa(in5)
	fmt.Println(s6)
	fmt.Printf("s6 = %T\n", s6)

	var h string = "Hello World"
	by := []byte(h)
	fmt.Println(by)
	h2 := string(by)
	fmt.Println(h2)
}
