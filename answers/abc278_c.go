// https://atcoder.jp/contests/adt_hard_20251216_1/tasks/abc278_c

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

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	ff := make([][]bool, N+1)
	for i := 1; i <= N; i++ {
		ff[i] = make([]bool, N+1)
	}

	for i := 0; i < Q; i++ {
		var T, A, B int
		fmt.Fscan(in, &T, &A, &B)
		if T == 1 {
			ff[A][B] = true
		}

		if T == 2 {
			ff[A][B] = false
		}

		if T == 3 {
			if ff[A][B] && ff[B][A] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
