// https://atcoder.jp/contests/adt_hard_20251218_1/tasks/abc381_d

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

	b := -1
	bCount := 0
	a := -1
	aCount := 0

	ans := 0
	
	for i := 0; i < N; i++ {
		var n int
		fmt.Fscan(in, &n)

		if b == -1 {
			b = n
			a = n
			bCount = 1
			continue
		}

		if b == a && b == n {
			bCount++
			continue
		}

		if b == a && a != n {
			a = n
			aCount = 1
			continue
		}

		if a == n {
			aCount++
			if bCount < aCount {
				t := bCount * 2
				if ans < t {
					ans = t
				}

				b = n
				bCount = aCount
			}
		} else {
			if bCount > aCount {
				t := aCount * 2
				if ans < t {
					ans = t
				}
			} else {
				t := bCount * 2
				if ans < t {
					ans = t
				}
			}

			b = a
			bCount = aCount
			a = n
			aCount = 1
		}
	}

	fmt.Fprintln(out, ans)
}
