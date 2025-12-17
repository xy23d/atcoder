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

	ns := []int{}
	a := make(map[int]bool)
	nIdx := make(map[int][]int)

	for i := 0; i < N; i++ {
		var n int
		fmt.Fscan(in, &n)
		if !a[n] {
			a[n] = true
			ns = append(ns, n)
		}

		nIdx[n] = append(nIdx[n], i)
	}

	ans := 0
	for _, n := range ns {
		t := execute(nIdx[n])

		if ans < t {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

func execute(ns []int) int {
	ans := 1

	for i := 0; i < len(ns); i++ {
		idx := ns[i]
		for j := i + 1; j < len(ns); j++ {
			idx2 := ns[j]

			d := idx2 - idx
			b := idx2
			t := 2

			for k := j+1; k < len(ns); k++ {
				idx3 := ns[k]

				if b + d < idx3 {
					break
				}

				if idx3 == b + d {
					b = idx3
					t++
				}
			}

			if ans < t {
				ans = t
			}
		}
	}

	return ans
}
