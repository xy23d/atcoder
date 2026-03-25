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

	var H, W int
	fmt.Fscan(in, &H, &W)
	as := make([][]int, H)
	for i := 0; i < H; i++ {
		as[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &as[i][j])
		}
	}

	var H2, W2 int
	fmt.Fscan(in, &H2, &W2)
	bs := make([][]int, H2)
	for i := 0; i < H2; i++ {
		bs[i] = make([]int, W2)
		for j := 0; j < W2; j++ {
			fmt.Fscan(in, &bs[i][j])
		}
	}

	ws := make([]int, W2)
	for j := 0; j < W2; j++ {
		ws[j] = -1
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if as[i][j] == bs[0][0] {
				ws[0] = j
				if check(as, bs, i, j+1, 0, 1, ws) {
					fmt.Fprintln(out, "Yes")
					return
				}
				ws[0] = -1
			}
		}
	}

	fmt.Fprintln(out, "No")
}

func check(as, bs [][]int, ax, ay, bx, by int, ws []int) bool {
	if len(bs) == bx {
		return true
	}

	if ws[len(bs[bx])-1] != -1 {
		for i := ax; i < len(as); i++ {
			r := true
			for j := 0; j < len(bs[bx]); j++ {
				if as[i][ws[j]] != bs[bx][j] {
					r = false
					break
				}
			}

			if r {
				ax++
				bx++
				if check(as, bs, i, ay, bx, by, ws) {
					return true
				}
				ax--
				bx--
			}
		}
	} else {
		for j := ay; j < len(as[ax]); j++ {
			if as[ax][j] == bs[bx][by] {
				ws[by] = j
				by++
				if len(bs[bx]) == by {
					ax++
					bx++
				}

				if check(as, bs, ax, j+1, bx, by, ws) {
					return true
				}
				if len(bs[bx]) == by {
					ax--
					bx--
				}
				by--
				ws[by] = -1
			}
		}
	}

	return false
}
