// https://atcoder.jp/contests/adt_hard_20260120_1/tasks/abc399_c

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

	uv := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		uv[i] = []int{}
	}

	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		uv[u] = append(uv[u], v)
		uv[v] = append(uv[v], u)
	}

	ans := 0
	a := make([]bool, N+1)
	for i := 1; i <= N; i++ {
		if a[i] {
			continue
		}

		a[i] = true
		ns := uv[i]

		nCounts := make([]int, N+1)
		for {
			tns := []int{}
			for _, n := range ns {
				a[n] = true
				nCounts[n]++

				if nCounts[n] > 1 {
					ans++
					continue
				}

				for _, n2 := range uv[n] {
					if a[n2] {
						continue
					}

					tns = append(tns, n2)
				}
			}

			if len(tns) == 0 {
				break
			}

			ns = tns
		}
	}

	fmt.Fprintln(out, ans)
}
