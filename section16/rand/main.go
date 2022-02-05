package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand
	rand.Seed(256)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int())
}
