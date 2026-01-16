// https://atcoder.jp/contests/adt_hard_20251218_1/tasks/abc395_c

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

	nIdx := make(map[int]int)

	ans := -1

	for i := 0; i < N; i++ {
		var n int
		fmt.Fscan(in, &n)

		if nIdx[n] != 0 {
			d := (i + 1 - nIdx[n]) + 1
			if ans == -1 || ans > d {
				ans = d
			}
		}

		nIdx[n] = i + 1
	}

	fmt.Fprintln(out, ans)
}
