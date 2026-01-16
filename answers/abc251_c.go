// https://atcoder.jp/contests/adt_hard_20251223_1/tasks/abc251_c

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

	a := make(map[string]bool)
	max := 0
	ans := 0
	for i := 0; i < N; i++ {
		var S string
		var T int
		fmt.Fscan(in, &S, &T)

		if a[S] {
			continue
		}

		a[S] = true

		if max < T {
			max = T
			ans = i + 1
		}
	}

	fmt.Fprintln(out, ans)
}
