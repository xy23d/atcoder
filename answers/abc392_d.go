// https://atcoder.jp/contests/adt_hard_20260122_1/tasks/abc392_d

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

	ks := make([]int, N)
	kns := make([][]int, N)
	kNCounts := make([]map[int]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ks[i])

		kns[i] = []int{}
		kNCounts[i] = make(map[int]int)
		a := make(map[int]bool)
		for j := 0; j < ks[i]; j++ {
			var n int
			fmt.Fscan(in, &n)

			if !a[n] {
				kns[i] = append(kns[i], n)
				a[n] = true
			}

			kNCounts[i][n]++
		}
	}

	ans := float64(0)
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			total := ks[i] * ks[j]

			count := 0
			for k := 0; k < len(kns[i]); k++ {
				n := kns[i][k]
				if kNCounts[i][n] > 0 && kNCounts[j][n] > 0 {
					count = count + kNCounts[i][n]*kNCounts[j][n]
				}
			}

			p := float64(count) / float64(total)
			if ans < p {
				ans = p
			}
		}
	}

	fmt.Fprintln(out, ans)
}
