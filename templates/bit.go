package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 0) // 1　左シフト
	fmt.Println(1 << 1) // 2 左シフト
	fmt.Println(3 & 1 << 1) // 2 アンド
	fmt.Println(3 | 1 << 1) // 3 OR
	fmt.Println(3 ^ 1 << 1) // 1 XOR
	fmt.Println((3>>2)&1 == 1) // そのbitに1がたっているかどうか
	fmt.Println(3^(1<<2)) // そのbitの1→0
	fmt.Println(3|(1<<2)) // そのbitに1をたてる
}
