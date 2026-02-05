// https://atcoder.jp/contests/adt_hard_20260205_1/tasks/abc273_c

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

	a := make(map[int]bool)
	ns := []int{}
	nCounts := make(map[int]int)

	for i := 0; i < N; i++ {
		var n int
		fmt.Fscan(in, &n)

		if !a[n] {
			a[n] = true
			ns = append(ns, n)
		}

		nCounts[n]++
	}

	sort.Ints(ns)
	for i := len(ns) - 1; i >= 0; i-- {
		fmt.Fprintln(out, nCounts[ns[i]])
	}

	for i := 0; i < N-len(ns); i++ {
		fmt.Fprintln(out, 0)
	}
}
