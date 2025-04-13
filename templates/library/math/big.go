package main

import (
	"fmt"
	"math/big"
)

func main() {
	// make([]*big.Int, N)

	calc = new(big.Int)
	
	n := big.NewInt(2)
	fmt.Println(calc.Add(n, big.NewInt(int64(1))))

	n = big.NewInt(2)
	fmt.Println(calc.Mul(n, big.NewInt(int64(2))))

	n = big.NewInt(2)
	fmt.Println(calc.Exp(n, big.NewInt(int64(2)), nil))
}
