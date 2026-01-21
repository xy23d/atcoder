// https://atcoder.jp/contests/adt_hard_20260121_1/tasks/abc413_c

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

	var Q int
	fmt.Fscan(in, &Q)

	xs := [][2]int{}
	idx := 0

	for i := 0; i < Q; i++ {
		var q int
		fmt.Fscan(in, &q)

		if q == 1 {
			var c, x int
			fmt.Fscan(in, &c, &x)

			xs = append(xs, [2]int{c, x})
		}

		if q == 2 {
			var k int
			fmt.Fscan(in, &k)

			ans := 0
			for j := idx; j < len(xs); j++ {
				c := xs[j][0]
				x := xs[j][1]

				if c > k {
					ans += x * k
					xs[j][0] = c - k
					fmt.Fprintln(out, ans)
					idx = j
					break
				} else {
					ans += c * x
				}
			}
		}
	}
}
