// https://atcoder.jp/contests/adt_medium_20250617_1/tasks/abc369_c

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(reader, &N)

	ns := make([]int, N)
	dss := make(map[int]map[int]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &ns[i])

		for j := 0; j <= i; j++ {
			if dss[j] == nil {
				dss[j] = make(map[int]int)
			}

			if i != j {
				dss[j][i] = (ns[i] - ns[j]) / (i - j)
			}
		}
	}

	r := 0
	for i := 0; i < N; i++ {
		ds := dss[i]
		count := 0
		b := 0
		for j := i + 1; j < N; j++ {
			if b == 0 {
				b = ds[j]
				count++
				continue
			}

			if b == ds[j] {
				count++
				continue
			}

			if count > 0 {
				if count == 1 {
					r = r + 1
				} else {
					r = r + factorial(count)/factorial(2)
				}

				count = 0
			}

			break
		}

		if count > 0 {
			if count == 1 {
				r = r + 1
			} else {
				r = r + factorial(count)/factorial(2)
			}
		}
	}

	fmt.Println(r)
}

// 階乗
func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
