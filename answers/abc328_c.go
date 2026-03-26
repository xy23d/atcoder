// https://atcoder.jp/contests/adt_all_20260326_1/tasks/abc328_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	var s string
	fmt.Fscan(in, &s)

	chars := strings.Split(s, "")
	b := ""
	counts := make([]int, N+1)
	counts[0] = 0
	for i := 0; i < N; i++ {
		counts[i+1] = counts[i]

		if b == chars[i] {
			counts[i+1]++
		}

		b = chars[i]
	}

	for i := 0; i < Q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, counts[r]-counts[l])
	}
}
