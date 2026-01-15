// https://atcoder.jp/contests/adt_hard_20260114_1/tasks/abc415_c

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

	var T int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		execute(in, out)
	}
}

func execute(in *bufio.Reader, out *bufio.Writer) {
	var N int
	fmt.Fscan(in, &N)

	var S string
	fmt.Fscan(in, &S)

	limit := 1 << N
	goal := limit - 1

	reachable := make([]bool, limit)
	reachable[0] = true

	for mask := 1; mask < limit; mask++ {
		if S[mask-1] == '1' {
			reachable[mask] = false
			continue
		}

		canReach := false
		for k := 0; k < N; k++ {
			if (mask>>k)&1 == 1 {
				prev := mask ^ (1 << k)

				if reachable[prev] {
					canReach = true
					break
				}
			}
		}

		if canReach {
			reachable[mask] = true
		}
	}

	if reachable[goal] {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
