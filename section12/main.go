package main

import "fmt"

// interface

// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)

}

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Number, c.Model)
}

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1}
}

func main() {
	// vs := []Stringfy{
	// 	&Person{Name: "Taro", Age: 21},
	// 	&Car{Number: "123-456", Model: "AB-1234"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

	// カスタムエラー
	// err := RaiseError()
	// fmt.Println(err)
	// // fmt.Println(err.Message)
	// e, ok := err.(*MyError)
	// if ok {
	// 	fmt.Println(e.Message)
	// }
	// fmt.Println(err.(*MyError).Message)

	// Stringer
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
