// https://atcoder.jp/contests/adt_hard_20251217_1/tasks/abc304_d

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

	var W, H, N int
	fmt.Fscan(in, &W, &H, &N)

	pq := make([][]bool, W+1)
	for i := 1; i <= W; i++ {
		pq[i] = make([]bool, H+1)
	}

	for i := 0; i < N; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		pq[p][q] = true
	}

	var A int
	fmt.Fscan(in, &A)
	as := make([]int, A)
	for i := 0; i < A; i++ {
		fmt.Fscan(in, &as[i])
	}

	var B int
	fmt.Fscan(in, &B)
	bs := make([]int, B)
	for i := 0; i < B; i++ {
		fmt.Fscan(in, &bs[i])
	}

	sort.Ints(as)
	sort.Ints(bs)

	min := N
	max := 0
	ba := 1
	bb := 1
	for i := 0; i < len(as); i++ {
		for j := 0; j < len(bs); j++ {
			t := 0
			for a := ba; a <= as[i]; a++ {
				for b := bb; b <= bs[j]; b++ {
					if pq[a][b] {
						t++
					}
				}
			}

			if min > t {
				min = t
			}

			if max < t {
				max = t
			}

			bb = bs[j]
		}

		ba = as[i]
		bb = 1
	}

	fmt.Fprintln(out, min, max)
}
