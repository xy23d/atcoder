// https://atcoder.jp/contests/adt_medium_20250702_1/tasks/abc337_d
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var H, W, K int
	fmt.Fscan(reader, &H, &W, &K)

	ss := make(map[int]map[int]string)
	ms := [][]int{}

	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(reader, &s)
		chars := strings.Split(s, "")
		ss[i] = make(map[int]string)
		for j := 0; j < W; j++ {
			ss[i][j] = chars[j]

			if ss[i][j] == "o" {
				ms = append(ms, []int{i, j})
			}
		}
	}

	mf := true
	var min int

	for _, vs := range ms {
		i := vs[0]
		j := vs[1]

		if i-(K-1) > -1 {
			k := (K - 1)
			r := 0
			for k > 0 {
				if ss[i+K-k][j] == "x" {
					break
				}

				if ss[i+K-k][j] == "." {
					r = r + 1
				}

				k = k - 1
			}

			if k == 0 {
				if mf || min > r {
					min = r
				}
			}
		}

		if i+(K-1) < H {
			k := (K - 1)
			r := 0
			for k > 0 {
				if ss[i+k-K][j] == "x" {
					break
				}

				if ss[i+k-K][j] == "." {
					r = r + 1
				}

				k = k - 1
			}

			if k == 0 {
				if mf || min > r {
					min = r
				}
			}
		}

		if j-(K-1) > -1 {
			k := (K - 1)
			r := 0
			for k > 0 {
				if ss[i][j+K-k] == "x" {
					break
				}

				if ss[i][j+K-k] == "." {
					r = r + 1
				}

				k = k - 1
			}

			if k == 0 {
				if mf || min > r {
					min = r
				}
			}
		}

		if j+(K-1) < W {
			k := (K - 1)
			r := 0
			for k > 0 {
				if ss[i][j+k-K] == "x" {
					break
				}

				if ss[i][j+k-K] == "." {
					r = r + 1
				}

				k = k - 1
			}

			if k == 0 {
				if mf || min > r {
					min = r
				}
			}
		}
	}

	fmt.Println(min)
}
