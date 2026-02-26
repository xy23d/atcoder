// https://atcoder.jp/contests/adt_all_20260226_1/tasks/abc315_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	fs := make(map[int][]int)
	a := make(map[int]bool)
	ns := []int{}
	for i := 0; i < N; i++ {
		var f, s int
		fmt.Fscan(in, &f, &s)

		if !a[f] {
			a[f] = true
			ns = append(ns, f)
			fs[f] = []int{}
		}

		fs[f] = append(fs[f], s)
	}

	ans := -1
	max1 := -1
	max2 := -1

	for i := 0; i < len(ns); i++ {
		f := ns[i]
		sort.Ints(fs[f])

		fLen := len(fs[f])
		if fLen >= 2 {
			t := fs[f][fLen-1] + fs[f][fLen-2]/2
			if ans < t {
				ans = t
			}
		}

		if max2 < fs[f][fLen-1] {
			max2 = fs[f][fLen-1]
			if max1 < max2 {
				max1, max2 = max2, max1
			}
		}
	}

	if ans < max1+max2 {
		ans = max1 + max2
	}

	fmt.Fprintln(out, ans)
}
