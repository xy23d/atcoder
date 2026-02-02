// https://atcoder.jp/contests/adt_hard_20260129_1/tasks/abc239_c

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

	var x1, y1, x2, y2 int
	fmt.Fscan(in, &x1, &y1, &x2, &y2)

	ns := []int{1, 2, -1, -2}
	dss := [][]int{}

	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(ns); j++ {
			if ns[i] == ns[j] || ns[i] == (-1) * ns[j] {
				continue
			}

			dss = append(dss, []int{ns[i], ns[j]})
		}
	}


	for _, ds := range dss {
		cx := x1 + ds[0]
		cy := y1 + ds[1]

		if (cx - x2) * (cx -x2) + (cy - y2) * (cy - y2) == 5 {
			fmt.Fprintln(out, "Yes")
			return
		}
	}

	fmt.Fprintln(out, "No")
}
