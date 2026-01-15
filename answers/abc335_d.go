// https://atcoder.jp/contests/adt_hard_20260113_1/tasks/abc335_d

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	n := 1
	ans := make([][]string, N)
	for i := 0; i < N; i++ {
		ans[i] = make([]string, N)
	}

	m := 0
	x := 0
	y := 0
	for {
		for {
			ans[x][y] = strconv.Itoa(n)
			n++

			if y+1 > N-1-m {
				break
			}

			y++
		}

		x++

		for {
			ans[x][y] = strconv.Itoa(n)
			n++

			if x+1 > N-1-m {
				break
			}

			x++
		}

		y--

		for {
			ans[x][y] = strconv.Itoa(n)
			n++

			if y-1 < 0+m {
				break
			}

			y--
		}

		x--

		for {
			ans[x][y] = strconv.Itoa(n)
			n++

			if x-1 < 0+1+m {
				break
			}

			x--
		}

		m++
		y++

		if x == (N - 1) / 2 && y == (N - 1) / 2 {
			ans[x][y] = "T"
			break
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fprintln(out, strings.Join(ans[i], " "))
	}
}
