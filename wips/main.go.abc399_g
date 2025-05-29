package main

import (
	"fmt"
)

func main() {
	var N, M , C int
	fmt.Scan(&N, &M, &C)

	as := make([]int, C)
	for i := 0; i < C; i++ {
		fmt.Scan(&as[i])
	}

	us := make([]int, M)
	vs := make([]int, M)
	cs := make([]int, M)
	css := make(map[int]int)
	for i := 0; i < M; i++ {
		fmt.Scan(&us[i])
		fmt.Scan(&vs[i])
		fmt.Scan(&cs[i])
		css[cs[i]]++
	}

	for i := 0; i < M; i++ {
	}

	result := 0

	for k := 0; k < C; k++ {
		if css[k] > as[k] {
			continue
		}

		for i := 0; i < C; i++ {
			L := i
			for j := i + 1; j < C; j++ {
				R := j

				if L <= cs[i] && cs[i] <= R {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}

