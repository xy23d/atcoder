// https://atcoder.jp/contests/adt_hard_20260114_1/tasks/abc256_c

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

	hs := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &hs[i])
	}

	ws := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &ws[i])
	}

	ans := 0

	for n00 := 1; n00 <= min(hs[0]-2, ws[0]-2); n00++ {
		for n01 := 1; n01 <= min(hs[1]-2, ws[0]-n00-1); n01++ {
			n02 := ws[0] - n00 - n01

			for n10 := 1; n10 <= min(hs[0]-n00-1, ws[1]-2); n10++ {
				n20 := hs[0] - n00 - n10

				for n11 := 1; n11 <= min(hs[1]-n01-1, ws[1]-n10-1); n11++ {
					n12 := ws[1] - n10 - n11
					n21 := hs[1] - n01 - n11
					n22 := hs[2] - n02 - n12

					if n20 > 0 && n21 > 0 && n22 > 0 &&  n20+n21+n22 == ws[2] {
						ans++
					}
				}
			}
		}
	}

	fmt.Fprintln(out, ans)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
