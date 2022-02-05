package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力を行単位で読み込む
	// 標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー: ", err)
	}

}
