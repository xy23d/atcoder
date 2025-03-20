package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ooo"
	chars := strings.Split(s, "")
	fmt.Println(chars)
	fmt.Println(strings.Join(chars, " "))
}
