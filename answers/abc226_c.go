// https://atcoder.jp/contests/adt_hard_20251217_1/tasks/abc226_c

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

	var N int
	fmt.Fscan(in, &N)

	as := make([][]int, N)
	ts := make([]int, N)

	for i := 0; i < N; i++ {
		var K int
		fmt.Fscan(in, &ts[i], &K)

		as[i] = []int{}

		for j := 0; j < K; j++ {
			var n int
			fmt.Fscan(in, &n)

			as[i] = append(as[i], n-1)
		}
	}

	ans := execute(N-1, ts[N-1], ts, as, make([]bool, N))

	fmt.Fprintln(out, ans)
}

func execute(n, total int, ts []int, as [][]int, a []bool) int {
	a[n] = true

	for _, v := range as[n] {
		if a[v] {
			continue
		}

		a[v] = true
		total = total + ts[v]
		total += execute(v, total, ts, as, a) - total
	}

	return total
}
