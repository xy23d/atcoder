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

			ans := int64(0)
			for {
				c := xs[idx][0]
				x := xs[idx][1]

				if c - k > 0 {
					ans += int64(k * x)
					xs[idx][0] = c - k
					break
				} else {
					ans += int64(c * x)
					k = k - c
					if k == 0 {
						idx++
						break
					}
				}

				idx++
			}

			fmt.Fprintln(out, ans)
		}
	}
}
