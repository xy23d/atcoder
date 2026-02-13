// https://atcoder.jp/contests/adt_hard_20260106_1/tasks/abc343_d

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

	var N, T int
	fmt.Fscan(in, &N, &T)

	ns := make([]int, N+1)
	nCount := make(map[int]int)
	nCount[0] = N
	ans := 1

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)

		n := ns[A]
		nCount[n]--
		if nCount[n] == 0 {
			ans--
		}

		ns[A] = n + B
		if nCount[ns[A]] == 0 {
			ans++
		}
		nCount[ns[A]]++

		fmt.Fprintln(out, ans)
	}
}
