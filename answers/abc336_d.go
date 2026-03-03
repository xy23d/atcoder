// https://atcoder.jp/contests/adt_hard_20260113_1/tasks/abc335_d

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

	ans := -1
	for i := 0; i < N; i++ {
		center := false
		tAns := -1
		t := -1

    as[i] = 1
		for j := i + 1; j < N; j++ {
			if !center {
				if as[j-1]+1 <= as[j] {
					continue
				}
				tAns = as[j] + 1
				t = tAns - 1
				if as[j] == 1 {
					break
				}
				center = true
			} else {
				if j+1 >= N {
					ans = -1
					break
				}

				if as[j]-1 <= as[j+1] {
					t--
					if t == 0 {
						break
					}
					continue
				}

				break
			}
		}

		if ans < tAns {
			ans = tAns
		}
	}

  if ans != -1 {
  	fmt.Fprintln(out, ans)
  } else {
    fmt.Fprintln(out, (N + 1) / 2)
  }
}
