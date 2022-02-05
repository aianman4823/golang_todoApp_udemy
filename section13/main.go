package main

import (
	"fmt"
	"section13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	return
}

func main() {
	fmt.Println(foo.Max)
	// foo.minは呼び出せないが、、、
	fmt.Println(foo.ReturnMin())

	// スコープ
	fmt.Println(appName())
	// fmt.Println(AppName)

	fmt.Println(Do("AAA"))
}
