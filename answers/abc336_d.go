// https://atcoder.jp/contests/adt_hard_20260113_1/tasks/abc335_d

package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	L := make([]int, N)
	R := make([]int, N)

	L[0] = min(A[0], 1)
	for i := 1; i < N; i++ {
		L[i] = min(A[i], L[i-1]+1)
	}

	R[N-1] = min(A[N-1], 1)
	for i := N - 2; i >= 0; i-- {
		R[i] = min(A[i], R[i+1]+1)
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans = max(ans, min(L[i], R[i]))
	}

	fmt.Fprintln(out, ans)
}
