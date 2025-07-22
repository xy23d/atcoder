package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(reader, &N)

	as := make([]int, N)
	bs := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &as[i])
	}

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &bs[i])
	}

	fmt.Println(calc(as[0], bs[0] - as[0] + 1, 1, as, bs, N))
}

func calc(min int, bCount int, index int, as []int, bs []int, N int) int {
	a := as[index]
	b := bs[index]
	max := bs[index - 1]

	var c int
	for i := min; i < max; i++ {
		c += i
	}

	fmt.Println("index", index)
	fmt.Println("c", c)

	count := (bCount * (b - a + 1)) - c % 998244353

	fmt.Println("count", count)

	if index == N - 1 {
		return count
	} else {
		return calc(count, index + 1, as, bs, N)
	}
}
