// https://atcoder.jp/contests/adt_hard_20260108_1/tasks/abc345_d

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

	var N, H, W int
	fmt.Fscan(in, &N, &H, &W)
	abs := make([][]int, N)

	for i := 0; i < N; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		abs[i] = []int{a, b}
	}

	exists := make([][]bool, H+1)
	for i := 1; i <= H; i++ {
		exists[i] = make([]bool, W+1)
	}

	a := make([]bool, N)
	for i := 0; i < N; i++ {
		if abs[i][0] <= H && abs[i][1] <= W {
			r := execute(i, 0, 0, abs[i][0], abs[i][1], N, H, W, abs, a, exists)
			if r {
				fmt.Fprintln(out, "Yes")
				return
			}
		}

		if abs[i][0] != abs[i][1] {
			if abs[i][1] <= H && abs[i][0] <= W {
				r := execute(i, 0, 0, abs[i][1], abs[i][0], N, H, W, abs, a, exists)
				if r {
					fmt.Fprintln(out, "Yes")
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "No")
}

func execute(idx, h, w, mh, mw, N, H, W int, abs [][]int, a []bool, exists [][]bool) bool {
	a[idx] = true

	for mh > 0{
		tMw := mw
		for tMw > 0 {
			exists[h+mh][w+tMw] = true
			tMw--
		}
		mh--
	}

	w = w + mw
	for w == W{
		h = h + 1
		w = 0
		for w < W {
			if !exists[h][w+1] {
				break
			}
			w++
		}

		if h == H && w == W {
			return true
		}
	}

	for i := 0; i < N; i++ {
		if a[i] {
			continue
		}

		if h+abs[i][0] <= H && w+abs[i][1] <= W {
			r := execute(i, h, w, abs[i][0], abs[i][1], N, H, W, abs, cloneABool(a), cloneAABool(exists))

			if r {
				return true
			}
		}

		if h+abs[i][1] <= H && w+abs[i][0] <= W {
			r := execute(i, h, w, abs[i][1], abs[i][0], N, H, W, abs, cloneABool(a), cloneAABool(exists))

			if r {
				return true
			}
		}
	}
	return false
}

func cloneABool(bs []bool) []bool {
	c := make([]bool, len(bs))

	for i := 0; i < len(bs); i++ {
		c[i] = bs[i]
	}

	return c
}

func cloneAABool(bss [][]bool) [][]bool {
	c := make([][]bool, len(bss))

	for i := 0; i < len(bss); i++ {
		c[i] = make([]bool, len(bss[i]))

		for j := 0; j < len(bss[i]); j++ {
			c[i][j] = bss[i][j]
		}
	}

	return c
}
