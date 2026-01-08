// https://atcoder.jp/contests/adt_hard_20260108_1/tasks/abc361_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K int
	fmt.Fscan(in, &N, &K)

	if N-K == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	as := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
	}

	sort.Ints(as)

	ans := 0
	for i := 0; i < N-1; i++ {
		idx := i + (N - K - 1)

		if idx >= N {
			break
		}

		t := abs(as[i] - as[idx])
		if ans == 0 || ans > t {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

func abs(n int) int {
	if n < 0 {
		return (-1) * n
	}

	return n
}
