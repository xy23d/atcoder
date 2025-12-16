// https://atcoder.jp/contests/adt_hard_20251216_1/tasks/abc369_d

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

	as := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
	}

	ans := 0
	sCount := 0
	for i := 0; i < N; i++ {
		if i+1 >= N {
			if (i+sCount)%2 == 0 {
				ans = ans + as[i]
			} else {
				ans = ans + as[i] * 2
			}
			continue
		}

		if as[i] < as[i+1] {
			if (i+sCount)%2 == 0 {
				ans = ans + as[i]
			} else {
				sCount++
				continue
			}
		} else {
			if (i+sCount)%2 == 0 {
				ans = ans + as[i]
			} else {
				ans = ans + as[i] * 2
			}
		}
	}

	fmt.Fprintln(out, ans)
}
