// https://atcoder.jp/contests/adt_hard_20260113_1/tasks/abc364_c

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

	var N, X, Y int
	fmt.Fscan(in, &N, &X, &Y)

	as := make([]int, N)
	aTotal := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
		aTotal += as[i]
	}

	bs := make([]int, N)
	bTotal := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &bs[i])
		bTotal += bs[i]
	}

	min := N

	if aTotal > X {
		sort.Ints(as)
		t1 := 0
		for i := 0; i < N; i++ {
			t1 = t1 + as[N-1-i]

			if t1 > X {
				if min > i+1 {
					min = i + 1
				}

				break
			}
		}
	}

	if bTotal > Y {
		sort.Ints(bs)
		t2 := 0
		for i := 0; i < N; i++ {
			t2 = t2 + as[N-1-i]

			if t2 > Y {
				if min > i+1 {
					min = i + 1
				}

				break
			}
		}
	}
	fmt.Fprintln(out, min)
}
