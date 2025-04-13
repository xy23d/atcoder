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

	ns := []int64{1, 2, 3}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i] < ns[j]
	})
}
