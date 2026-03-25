// https://atcoder.jp/contests/adt_all_20260325_1/tasks/abc331_c

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

	as := make([]int, N)
	a := make(map[int]bool)
	ns := []int{}
	nCounts := make(map[int]int)
	total := 0

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])

		nCounts[as[i]]++
		if !a[as[i]] {
			ns = append(ns, as[i])
			a[as[i]] = true
		}

		total = total + as[i]
	}

	sort.Ints(ns)

	diff := 0
	diffs := make(map[int]int)
	for i := 0; i < len(ns); i++ {
		diff = diff + ns[i]*nCounts[ns[i]]
		diffs[ns[i]] = diff
	}

	for i := 0; i < N; i++ {
		if i != 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, total - diffs[as[i]])
	}

	fmt.Fprintln(out, "")
}
