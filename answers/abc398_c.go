// https://atcoder.jp/contests/adt_hard_20251218_1/tasks/abc395_c

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

	nToI := make(map[int][]int)
	a := make(map[int]bool)
	ns := []int{}

	for i := 0; i < N; i++ {
		var n int
		fmt.Fscan(in, &n)

		if !a[n] {
			ns = append(ns, n)
			a[n] = true
		}

		if nToI[n] == nil {
			nToI[n] = []int{}
		}

		nToI[n] = append(nToI[n], i+1)
	}

	sort.Ints(ns)
	for i := 0; i < len(ns); i++ {
		n := ns[len(ns)-1-i]
		if len(nToI[n]) == 1 {
			fmt.Fprintln(out, nToI[n][0])
			return
		}
	}

	fmt.Fprintln(out, -1)
}
