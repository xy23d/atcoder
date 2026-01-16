// https://atcoder.jp/contests/adt_hard_20260106_1/tasks/abc407

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

  b, _ := strconv.Atoi(chars[0])
	ans := b + 1
	for i := 1; i < len(chars); i++ {
     n, _ := strconv.Atoi(chars[i])
		if b < n {
			ans = ans + 10
		}

    ans++

		b = n
	}

	fmt.Fprintln(out, ans)
}
