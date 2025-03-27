package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s[0:]) // [2 3 4]
	fmt.Println(s[:1]) // [1]
	fmt.Println(append(s[:1], s[2:]...)) // [1 3 4]
}
