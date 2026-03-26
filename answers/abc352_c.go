// https://atcoder.jp/contests/adt_all_20260326_1/tasks/abc352_c

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

	abs := make([][2]int, N)
	total := 0
	max := -1

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &abs[i][0], &abs[i][1])
		total = total + abs[i][0]

		if max == -1 || max < abs[i][1]-abs[i][0] {
			max = abs[i][1] - abs[i][0]
		}
	}

	fmt.Fprintln(out, total+max)
}
