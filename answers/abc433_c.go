// https://atcoder.jp/contests/adt_hard_20260106_1/tasks/abc433_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var S string
	fmt.Fscan(in, &S)

	chars := strings.Split(S, "")
	ans := 0
	for i := 0; i < len(chars) - 1; i++ {
		s := chars[i]

		j := i+1
		sameCount := 1
		for j < len(chars) {
			if s != chars[j] {
				break
			}

			j++
			sameCount++
		}

		if j >= len(chars) - 1 {
			continue
		}

		c := chars[j]

		n1, _ := strconv.Atoi(s)
		n2, _ := strconv.Atoi(c)

		if n1 + 1 != n2 {
			continue
		}

		j++
		sameCount--
		for sameCount > 0 && j < len(chars) {
			if c != chars[j] {
				break
			}

			sameCount--
			j++
		}

		if sameCount == 0 {
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}
