// https://atcoder.jp/contests/adt_hard_20251217_1/tasks/abc385_c

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

	ns := make([]int, N)
	nUnique := []int{}
	a := make(map[int]bool)
	nIdx := make(map[int][]int)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ns[i])
		n := ns[i]
		if !a[n] {
			a[n] = true
			nUnique = append(nUnique, n)
		}

		nIdx[n] = append(nIdx[n], i)
	}

	ans := 0
	for _, n := range nUnique {
		t := execute(n, N, ns, nIdx[n])

		if ans < t {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

func execute(n, N int, ns, nIdx []int) int {
	ans := 1

	for i := 0; i < len(nIdx); i++ {
		idx := nIdx[i]
		for j := i + 1; j < len(nIdx); j++ {
			idx2 := nIdx[j]

			d := idx2 - idx
			t := 2

			for k := idx2 + d; k < N; k = k + d {
				if ns[k] != n {
					break
				}

				t++
			}

			if ans < t {
				ans = t
			}
		}
	}

	return ans
}
