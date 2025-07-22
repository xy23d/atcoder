package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	ch := make(map[int]int)
	cs := [][]int

	for i := 0; i< N; i++ {
		ch[i] = i
		cs = append(cs, []int{i, i})
	}

	for i := 0; i < Q; i++ {
		var n int
		fmt.Fscan(reader, &n)

		if n == 1 {
			var a, b int
			fmt.Fscan(reader, &a, &b)
			
			ch[a] = b
		}

		if n == 2 {
			var a, b int
			fmt.Fscan(reader, &a, &b)

			cs = append(cs, []int{a, b})
		}

		if n == 3 {
			var a int
			fmt.Fscan(reader, &a)
			c := hs[a]
		}
	}
}
