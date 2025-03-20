package main

import (
	"fmt"
	"math/big"
)

func main() { 
	n := big.NewInt(2)
	fmt.Println(n.Add(n, big.NewInt(int64(1))))

	n = big.NewInt(2)
	fmt.Println(n.Mul(n, big.NewInt(int64(2))))

	n = big.NewInt(2)
	fmt.Println(n.Exp(n, big.NewInt(int64(2)), nil))
}