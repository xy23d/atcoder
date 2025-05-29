package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	total := 0
	mins := make(map[int]int)

	for i := 0; i < N; i++ {
		var A, B, X int

		fmt.Scan(&A, &B, &X)

		if mins[i+1] != 0 {
			if mins[i+1] < total {
				total = mins[i+1]
			}
		}

		if mins[X] != 0 {
			if mins[X] > total + B {
				mins[X] = total + B
			}
		} else {
			mins[X] = total + B
		}

		total += A
	}

	result := 0

	fmt.Println(total)
}