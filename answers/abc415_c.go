// https://atcoder.jp/contests/adt_hard_20260114_1/tasks/abc415_c

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

	var T int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		execute(in, out)
	}
}

func execute(in *bufio.Reader, out *bufio.Writer) {
	var N int
	fmt.Fscan(in, &N)

	var S int
	fmt.Fscan(in, &S)

	a := make([]bool, N+1)
	if subExecute(N, 0, a, S) {
		fmt.Println(out, "Yes")
	} else {
		fmt.Println(out, "No")
	}
}

func subExecute(N, total int, a []bool, S byte) bool {
	for i := 1; i < N; i++ {
		if a[i] {
			continue
		}

		if (total+i)<<1&S > 1 {
			continue
		}

		a[i] = true

		if check(a) {
			return true
		}

		r := subExecute(N, total+i, a, S)
		if r {
			return true
		}

		a[i] = false
	}

	return false
}

func check(N int, a []bool) bool {
	for i := 1; i <= N; i++ {
		if !a[i] {
			return false
		}
	}

	return true
}
