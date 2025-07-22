package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	As := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&As[i])
	}

	a := big.NewInt(0)
	for i := 0; i < len(As); i++ {
		t := big.NewInt(int64(As[i]))
		a.Add(a, t.Exp(t, big.NewInt(int64(K)), nil))
	}

	b := big.NewInt(0)
	for i := 0; i < len(As); i++ {
		t := As[i]

		for j := i + 1; j < len(As); j++ {
			t += As[j]
			t2 := big.NewInt(int64(t))
			b.Add(b, t2.Exp(t2, big.NewInt(int64(K)), nil))
		}
	}

	t := a.Add(a, b)

	fmt.Println(t.Mod(t, big.NewInt(int64(998244353))))
}
