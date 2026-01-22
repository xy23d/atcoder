// https://atcoder.jp/contests/adt_hard_20260122_1/tasks/abc245_c

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

	var N, K int
	fmt.Fscan(in, &N, &K)

	as := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
	}

	bs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &bs[i])
	}

	ns := []int{as[0], bs[0]}
	for i := 1; i < N; i++ {
		tns := []int{}

		a := false
		b := false

		if as[i] == bs[i] {
			b = true
		}

		for j := 0; j < len(ns); j++ {
			if !a {
				if abs(ns[j]-as[i]) <= K {
					a = true
					tns = append(tns, as[i])
				}
			}

			if !b {
				if abs(ns[j]-bs[i]) <= K {
					b = true
					tns = append(tns, bs[i])
				}
			}

			if a && b {
				break
			}
		}

		if len(tns) == 0 {
			fmt.Fprintln(out, "No")
			return
		}

		ns = tns
	}

	fmt.Fprintln(out, "Yes")
}

func abs(n int) int {
	if n < 0 {
		return (-1) * n
	}

	return n
}
