// https://atcoder.jp/contests/adt_hard_20260203_1/tasks/abc389_c

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

	dTotal := 0
	total := 0
	idx := 0
	ns := []int{}
	ns = append(ns, 0)

	for i := 0; i < N; i++ {
		var q int
		fmt.Fscan(in, &q)

		if q == 1 {
			var l int
			fmt.Fscan(in, &l)

			total = total + l
			ns = append(ns, total)
		}

		if q == 2 {
			dTotal = ns[idx+1]
			idx++
		}

		if q == 3 {
			var k int
			fmt.Fscan(in, &k)
			fmt.Fprintln(out, ns[idx+k-1]-dTotal)
		}
	}
}
