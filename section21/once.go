package main

import (
	"fmt"
	"sync"
)

var onceA, onceB sync.Once

func A() {
	fmt.Println("A")
	onceB.Do(B)
}

func B() {
	fmt.Println("B")
	onceA.Do(A)
}

func main() {
	A()
}
