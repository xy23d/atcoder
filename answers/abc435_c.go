// https://atcoder.jp/contests/adt_hard_20260121_1/tasks/abc435_c

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
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ns[i])
	}

	max := 1
	for i := 0; i < max; i++{
		if max < (ns[i] + i) {
			max = ns[i] + i

			if max >= N {
				max = N
				break
			}
		}
	}

	fmt.Fprintln(out, max)
}
