package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var S string
	fmt.Fscan(in, &N, &S)

	chars := strings.Split(S, "")

	nIdx := make([]int, N+1)
	nl := 
	ns[0] = 0
	idx := 0


	for i := 1; i <= N; i++ {
		if chars[i-1] == "L" {
			t := ns[idx]
			ns[i] = i - 1
			ns[i-1] = i - 2
		}

		if chars[i-1] == "R" {
			t := ns[i]
			ns[i] = i - 1
			ns[i-1] = t
		}
	}

	fmt.Fprint(out, f, ns)

// 	fmt.Fprint(out, f)

// 	n := ns[f]
// 	for n != -1 {
// 		fmt.Fprint(out, " ", n)
// 		n = ns[n]
// 	}

// 	fmt.Fprintln(out)
}
