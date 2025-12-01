package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println(1 << 0) // 1
	fmt.Println(1 << 1) // 2
	fmt.Println(3 & 1 << 1) // 2
	fmt.Println(3 | 1 << 1) // 3

	var N, T, M int
	fmt.Fscan(in, &N, &T, &M)

	ng := make([]int, N)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		ng[a] |= 1<<b
		ng[b] |= 1<<a
	}

	limit := 1 << N
	ok := make([]bool, limit)
	ok[0] = true
	for mask := 1; mask < limit; mask++ {
		valid := true
		for i := 0; i < N && valid; i++ {
			if (mask>>i)&1 != 0 {
				if (ng[i] & mask) != 0 {
					valid = false
					break
				}
			}
		}

		ok[mask] = valid
	}

	dp := make([][]int64, limit)
	for i := range dp {
		dp[i] = make([]int64, T+1)
	}
	dp[0][0] = 1
	for mask := 1; mask < limit; mask++ {
		lsb := mask & -mask
		rest := mask ^ lsb

		for s := rest; ; s = (s - 1) & rest {
			group := s | lsb
			if ok[group] {
				for k := 1; k <= T; k++ {
					dp[mask][k] += dp[mask^group][k-1]
				}
			}

			if s == 0 {
				break
			}
		}
	}

	fmt.Fprintln(out, dp[limit-1][T])
}

func bitIndex(x int) int {
	idx := 0
	for (x & 1) == 0 {
		x >>= 1
		idx++
	}
	return idx
}

