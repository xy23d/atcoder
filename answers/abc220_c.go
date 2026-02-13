// https://atcoder.jp/contests/adt_hard_20251218_1/tasks/abc220_c

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

	var N int
	fmt.Fscan(in, &N)

	as := make([]int, N)
	total := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
		total = total + as[i]
	}

	var X int
	fmt.Fscan(in, &X)

	d := (X / total)

	diff := X - total*d

	ans := N * d

	for i := 0; i < N; i++ {
		ans++

		if diff < as[i] {
			fmt.Fprintln(out, ans)
			return
		}

		diff = diff - as[i]
	}
}
