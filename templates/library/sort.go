package main

import (
	"fmt"
	"sort"
)

func main() {
	ns := []int{3, 2, 1}
	sort.Ints(ns)
	fmt.Println(ns)

	ss := []string{"c", "b", "a"}
	sort.Strings(ss)
	fmt.Println(ss)
}
