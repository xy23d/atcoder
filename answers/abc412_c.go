// https://atcoder.jp/contests/adt_hard_20251125_1/tasks/abc412_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var start int64
	var end int64

	ss := make([]int64, N-1)
	for i := 0; i < N; i++ {
		var n int64
		fmt.Fscan(in, &n)

		if i == 0 {
			start = n
			continue
		}

		if i == N-1 {
			end = n
			continue
		}

		ss[i-1] = n
	}

	current := start * 2
	ans := 2

	if current >= end {
		fmt.Fprintln(out, ans)
		return
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	idx := -1
	for {
		exits := false

		for i := N - 1 - 1; i > idx; i-- {
			if current >= ss[i] {
				ans++

				if ss[i]*2 >= end {
					fmt.Fprintln(out, ans)
					return
				}

				exits = true
				current = ss[i] * 2
				idx = i
				break
			}
		}

		if !exits {
			break
		}
	}

	fmt.Fprintln(out, -1)
}
