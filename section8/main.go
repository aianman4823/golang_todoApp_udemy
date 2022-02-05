package main

import (
	"fmt"
	"strconv"
	"time"
)

// 条件分岐
func main() {
	a := 1
	if a == 2 {
		fmt.Println("Two")

	} else if a == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("I don't know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	// fmt.Println(x)

	// エラーハンドリング
	var s string = "A"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	// for
	/*ifor := 0
	for {
		ifor++
		if ifor == 3 {
			break
		}
		fmt.Println("loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i3 := 0; i3 < 10; i3++ {
		if i3 == 3 {
			continue
		}
		fmt.Println(i3)
	}
	*/

	// arr := [3]int{1, 2, 3}
	// for i2 := 0; i2 < len(arr); i2++ {
	// 	fmt.Println(i2)
	// }

	// arr := [3]int{1, 2, 3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// n := 3
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }

	// 型スイッチ
	anything("aaa")
	anything(1)

	var x interface{} = 3
	// 型アサーション
	i3, isInt := x.(int)
	fmt.Println(i3+2, isInt)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("I don't know")
	}

	// ラベルつきfor
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("Start")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")

	// Loop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 3; j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i, j, i*j)
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}

	// defer
	TestDefer()
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	RunDefer()

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// file.Write([]byte("Hello"))

	// panic recover
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("Start")

	// go goroutin
	// go sub()
	// go sub()
	// for {
	// 	fmt.Println("Main loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// init関数はMain関数よりも先に呼ばれる
	fmt.Println("Main")

}

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 1000)
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("Start")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init2")
}
