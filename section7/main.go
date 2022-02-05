package main

import (
	"fmt"
)

// 関数
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i6)

	// 関数を返す関数
	f2 := ReturnFunc()
	f2()

	// 関数を引数に取る関数
	CallFunction(func() { fmt.Println("I'm also function") })

	// クロージャ
	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Golang"))

	// ジェネレータ(クロージャを利用すると作れる)
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints:= integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

}
