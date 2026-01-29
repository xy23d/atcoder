// https://atcoder.jp/contests/adt_hard_20260129_1/tasks/abc239_c

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

	var x1, y1, x2, y2 int
	fmt.Fscan(in, &x1, &y1, &x2, &y2)

	if x1 == x2 && y1 == y2 {
		fmt.Fprintln(out, "Yes")
		return
	}

	if x1 == y2 && y1 == x2 {
		fmt.Fprintln(out, "Yes")
		return
	}

	nss := []int{}
	nss = append(nss, [1, 2])
	nss = append(nss, [2, 1])

	for _, ns := range nss {
		n1 := nss[0]
		n2 := nss[1]

		if x1 + n2 == y2 - n1 && y1 - n1
	}

	if x1+2 == y2-1 && y1+1 == x2-2 {
		fmt.Fprintln(out, "Yes")
		return
	}

	if x1-2 == y2+1 && y1-1 == x2+2 {
		fmt.Fprintln(out, "Yes")
		return
	}

	if x1+1 == y2-2 && y1+2 == x2-1 {
		fmt.Fprintln(out, "Yes")
		return
	}

	if x1-1 == y2+2 && y1-2 == x2+1 {
		fmt.Fprintln(out, "Yes")
		return
	}

	fmt.Fprintln(out, "No")
}
