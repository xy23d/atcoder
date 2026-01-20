// https://atcoder.jp/contests/adt_hard_20260120_1/tasks/abc288_c

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

	var N, M int
	fmt.Fscan(in, &N, &M)

	ns := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		ns[a] = append(ns[a], b)
		ns[b] = append(ns[b], a)
	}

	ans := 0
	a := make([]bool, N+1)
	for i := 1; i <= N; i++ {
		if a[i] {
			continue
		}

		a[i] = true

		nCounts := make([]int, N+1)
		nCounts[i]++

		tns := [][]int{}
		for _, n := range ns[i] {
			tns = append(tns, []int{n, i})
		}

		for i := 0; i < len(tns); i++ {
			n := tns[i][0]
			b := tns[i][1]

			a[n] = true

			nCounts[n]++

			if nCounts[n] > 1 {
				ans++
				continue
			}

			for _, n2 := range ns[n] {
				if n2 == b {
					continue
				}

				tns = append(tns, []int{n2, n})
			}
		}
	}

	fmt.Fprintln(out, ans / 2)
}
