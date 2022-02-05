package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 文字列に含まれる部分文字列を検索する
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

}
