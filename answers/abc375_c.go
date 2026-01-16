// https://atcoder.jp/contests/adt_hard_20260113_1/tasks/abc375_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	ass := make([][]string, N+1)
	for i := 1; i <= N; i++ {
		ass[i] = make([]string, N+1)

		var as string
		fmt.Fscan(in, &as)
		chars := strings.Split(as, "")

		for j := 1; j <= N; j++ {
			ass[i][j] = chars[j-1]
		}
	}

	for i := 1; i <= N/2; i++ {
		tAss := cloneStrings(ass)
		for x := i; x <= N+1-i; x++ {
			for y := i; y <= N+1-i; y++ {
				tAss[y][N+1-x] = ass[x][y]
			}
		}
		ass = tAss
	}

	for x := 1; x <= N; x++ {
		fmt.Fprintln(out, strings.Join(ass[x], ""))
	}
}

func cloneStrings(ass [][]string) [][]string {
	r := make([][]string, len(ass))

	for i := 0; i < len(ass); i++ {
		r[i] = make([]string, len(ass[i]))
		for j := 0; j < len(ass[i]); j++ {
			r[i][j] = ass[i][j]
		}
	}

	return r
}
