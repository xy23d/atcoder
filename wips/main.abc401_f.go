package main

import (
	"fmt"
)

func main() {
	var N1 int
	fmt.Scan(&N1)

	vu1 := make(map[int][]int)
	uv1 := make(map[int][]int)
	for i := 0; i < N1 - 1; i++ {
		var v, u int
		fmt.Scan(&v, &u)
	
		if vu1[u] == nil {
			vu1[u] = []int{}
		}
		if uv1[v] == nil {
			uv1[v] = []int{}
		}
		vu1[u] = append(vu1[u], v)
		uv1[v] = append(uv1[v], u)
	}

	var N2 int
	fmt.Scan(&N2)

	vu2 := make(map[int][]int)
	uv2 := make(map[int][]int)
	for i := 0; i < N1 - 1; i++ {
		var v, u int
		fmt.Scan(&v, &u)
	
		if vu2[u] == nil {
			vu2[u] = []int{}
		}
		if uv2[v] == nil {
			uv2[v] = []int{}
		}
		vu2[u] = append(vu2[u], v)
		uv2[v] = append(uv2[v], u)
	}

	result := 0
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			
		}
	}
	
	fmt.Println(result)
}

func n(parent int, i int, vu map[int][]int, uv map[int][]int) int {
	r := 0

	count1 := 0
	for i := 0; i < vu[i]; i++ {
		if vu[i] == parent {
			continue
		}

		if len(vu[vu[i]]) == 0 {
			break
		}
	}
}
