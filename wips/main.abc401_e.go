package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	vu := make(map[int][]int)
	ns := make(map[int]bool)
	for i := 0; i < N; i++ {
		var v, u int
		fmt.Scan(&v, &u)

		if vu[v] == nil {
			vu[v] = []int{}
		}
		vu[v] = append(vu[v], u)
		ns[v] = true
		ns[u] = true
	}

	fmt.Println(len(vu[1]))

	for i := 2; i <= M; i++ {
		result := 0
	
		r := false
		for j := 1; j <= i; j++ {
			for k := 0; k < len(vu[j]); k++ {
				if vu[j][k] == j {
					r = true
				} else {
					result++
				}
			}
		}

		if r {
			fmt.Println(result)
		} else {
			fmt.Println(-1)
		}
	}
}
