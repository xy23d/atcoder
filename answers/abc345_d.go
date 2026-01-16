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
	abs := make([][2]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &abs[i][0], &abs[i][1])
	}

	exists := make([][]bool, H)
	for i := 0; i < H; i++ {
		exists[i] = make([]bool, W)
	}

	al := make([]bool, N)
	for i := 1; i < 2; i++ {
		a := abs[i][0] - 1
		b := abs[i][1] - 1

		if a <= H && b <= W {
			r := execute(i, 0, 0, a, b, N, H, W, abs, al, cloneAABool(exists))
			if r {
				fmt.Fprintln(out, "Yes")
				return
			}
		}

		if a != b {
			if b <= H && a <= W {
				r := execute(i, 0, 0, b, a, N, H, W, abs, al, cloneAABool(exists))
				if r {
					fmt.Fprintln(out, "Yes")
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "No")
}

func execute(idx, h, w, mh, mw, N, H, W int, abs [][2]int, al []bool, exists [][]bool) bool {
	al[idx] = true

	for mh >= 0 {
		tMw := mw
		for tMw >= 0 {
			if exists[h+mh][w+tMw] {
				return false
			}

			exists[h+mh][w+tMw] = true

			tMw--
		}
		mh--
	}

	w = w + mw

	if w == W-1 {
		w++
		for w == W {
			w = 0
			h = h + 1
			for w < W {
				if !exists[h][w] {
					break
				}
				w++
			}

			if h == H-1 && w == W {
				return true
			}
		}
	} else {
		w = w + 1
	}

	for i := 0; i < N; i++ {
		if al[i] {
			continue
		}

		a := abs[i][0] - 1
		b := abs[i][1] - 1
		if h+a < H && w+b < W {
			r := execute(i, h, w, a, b, N, H, W, abs, al, cloneAABool(exists))

			if r {
				return true
			}
		}

		if a != b {
			if h+b < H && w+a < W {
				r := execute(i, h, w, b, a, N, H, W, abs, al, cloneAABool(exists))

				if r {
					return true
				}
			}
		}
	}

	al[idx] = false

	return false
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
