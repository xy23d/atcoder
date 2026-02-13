package main

// 問題文
// 整数 
// N が与えられるので、以下の条件を全て満たす整数 
// X の個数を 
// 998244353 で割った余りを求めてください。

// N 桁の正整数である。
// X の各桁を上の位から順に 
// X 
// 1
// ​
//  ,X 
// 2
// ​
//  ,…,X 
// N
// ​
//   とする。このとき以下の条件を全て満たす。
// 全ての整数 
// 1≤i≤N に対して、 
// 1≤X 
// i
// ​
//  ≤9
// 全ての整数 
// 1≤i≤N−1 に対して、 
// ∣X 
// i
// ​
//  −X 
// i+1
// ​
//  ∣≤1
// 制約
// N は整数
// 2≤N≤10 
// 6

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

	dp := make([][]int, N)
	dp[0] = make([]int, 10)
	for i := 1; i <= 9; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < N; i++ {
		dp[i] = make([]int, 10)

		for j := 1; j <= 9; j++ {
			for k := 1; k <= 9; k++ {
				if abs(j-k) <= 1 {
					dp[i][k] = (dp[i][k] + dp[i-1][j]) % 998244353
				}
			}
		}
	}

	ans := 0
	for _, v := range dp[N-1] {
		ans = (ans + v) % 998244353
	}

	fmt.Fprintln(out, ans)
}

func abs(n int) int {
	if n < 0 {
		return (-1) * n
	}

	return n
}
