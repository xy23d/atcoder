// https://atcoder.jp/contests/adt_hard_20251223_1/tasks/abc398_d

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

	var N, R, C int
	fmt.Fscan(in, &N, &R, &C)

	ss := make([][]int, N)
	ss[0] = []int{0, 0}

	var S string
	fmt.Fscan(in, &S)

	chars := strings.Split(S, "")

	ans := []string{}
	for i := 0; i < N; i++ {
		s := chars[i]

		var mi int
		var mj int

		if s == "N" {
			mi = -1
		}

		if s == "S" {
			mi = 1
		}

		if s == "W" {
			mj = -1
		}

		if s == "E" {
			mj = 1
		}

		r := false
		for j := 0; j <= i; j++ {
			ss[j][0] = ss[j][0] + mi
			ss[j][1] = ss[j][1] + mj

			if ss[j][0] == R && ss[j][1] == C {
				r = true
			}
		}

		if R == 0 && C == 0 {
			r = true
		}

		if i+1 < N {
			ss[i+1] = []int{0, 0}
		}

		if r {
			ans = append(ans, "1")
		} else {
			ans = append(ans, "0")
		}
	}

	fmt.Fprintln(out, strings.Join(ans, ""))
}
