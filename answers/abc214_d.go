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

	uvw := make([][]int, N+1)
	for u := 1; u <= N; u++ {
		uvw[u] = make([]int, N+1)
	}

	for i := 0; i < N-1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)

		if uvw[u][v] < w {
			uvw[u][v] = w
			uvw[v][u] = w
		}
	}

	for u := 1; u <= N; u++ {
		ns := make
		for v, w := range uvw[u] {

		}
	}

	fmt.Println(uvw)

	ans := 0
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			ans = ans + uvw[i][j]
		}
	}

	fmt.Fprintln(out, ans)
}
