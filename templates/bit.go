package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 0) // 1
	fmt.Println(1 << 1) // 2
	fmt.Println(3 & 1 << 1) // 2
	fmt.Println(3 | 1 << 1) // 3
}
