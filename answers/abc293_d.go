// https://atcoder.jp/contests/adt_hard_20260204_1/tasks/abc293_d

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

	var n, m int
	fmt.Fscan(in, &n, &m)
	rbs := make([][2]int, n+1)

	for i := 0; i < m; i++ {
		var n1, n2 int
		var color1, color2 string
		fmt.Fscan(in, &n1, &color1, &n2, &color2)

		if color1 == "R" {
			rbs[n1][0] = n2
		} else {
			rbs[n1][1] = n2
		}

		if color2 == "R" {
			rbs[n2][0] = n1
		} else {
			rbs[n2][1] = n1
		}
	}

	a := make([]bool, n+1)
	ans1 := 0
	ans2 := 0
	for i := 1; i <= n; i++ {
		if a[i] {
			continue
		}

		a[i] = true
		if execute(n, i, rbs, a) {
			ans1++
		} else {
			ans2++
		}
	}

	fmt.Fprintln(out, ans1, ans2)
}

func execute(n, an int, rbs [][2]int, a []bool) bool {
	ns2 := []int{}
	ns2 = append(ns2, an)

	ans := true
	for i := 0; i < len(ns2); i++ {
		n2 := ns2[i]
		for _, n3 := range rbs[n2] {
			if n3 == 0 {
				ans = false
			}

			if !a[n3] {
				a[n3] = true
				ns2 = append(ns2, n3)
			}
		}
	}

	return ans
}
