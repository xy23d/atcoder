// https://atcoder.jp/contests/adt_hard_20260205_1/tasks/abc332_c

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

	var N, M int
	fmt.Fscan(in, &N, &M)

	var s string
	fmt.Fscan(in, &s)

	chars := strings.Split(s, "")

	max := 0
	t := 0
	for i := 0; i < N; i++ {
		if chars[i] == "2" {
			t++
		}

		if chars[i] == "0" {
			if max < t {
				max = t
			}
		}
	}

	if max < t {
		max = t
	}

	t = max + M
	for i := 0; i < N; i++ {
		if chars[i] == "1" || chars[i] == "2" {
			t--
			if t < 0 {
				max++
			}
		}

		if chars[i] == "0" {
			t = max + M
		}
	}

	fmt.Fprintln(out, max)
}
