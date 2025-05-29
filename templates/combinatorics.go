package main

import "fmt"

// 階乗
func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

// 順列: nPr = n! / (n - r)!
func nPr(n int, r int) int {
	return factorial(n) / factorial(n - r)
}

// 組合せ: nCr = nPr / r!
func nCr(n int, r int) int {
	return nPr(n, r) / factorial(r)
}