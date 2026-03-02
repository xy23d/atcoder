// https://atcoder.jp/contests/adt_hard_20260108_1/tasks/abc423_c

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

	var N, R int
	fmt.Fscan(in, &N, &R)

	ls := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ls[i])
	}

	ans := 0
	is1 := false
	for i := 0; i < R-1; i++ {
		if !is1 {
			if ls[i] == 1 {
				continue
			}
		}

		if ls[i] == 1 {
			ans = ans + 2
		} else {
			is1 = true
			ans = ans + 1
		}
	}

	is2 := false
	for i := N-1; i >= R-1; i-- {
		if !is2 {
			if ls[i] == 1 {
				continue
			}
		}

		if ls[i] == 1 {
			ans = ans + 2
		} else {
			is2 = true
			ans = ans + 1
		}
	}

	fmt.Fprintln(out, ans)
}
