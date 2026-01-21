// https://atcoder.jp/contests/adt_hard_20260121_1/tasks/abc435_c

// https://atcoder.jp/contests/adt_hard_20251125_1/tasks/abc412_c

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

	ns := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ns[i])
	}

	ans := 1
	nsB := make([]bool, N)
	nsB[0] = true

	for i := 0; i < N; i++ {
		if !nsB[i] {
			break
		}

		for j := i+1; j < min(i + ns[i], N); j++{
			if !nsB[j] {
				nsB[j] = true
				ans++
			}
		}
	}

	fmt.Fprintln(out, ans)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
