// https://atcoder.jp/contests/adt_hard_20251211_1/tasks/abc239_d

// 問題文
// 高橋君と青木君が次のようなゲームをします。

// まず、高橋君が 
// A 以上 
// B 以下の好きな整数を選び、青木君に伝える
// 次に、青木君が 
// C 以上 
// D 以下の好きな整数を選ぶ
// 二人の選んだ整数の和が素数なら青木君の勝ち、そうでなければ高橋君の勝ち
// 二人が最適な戦略を取るとき、どちらが勝ちますか？

// 制約
// 1≤A≤B≤100
// 1≤C≤D≤100
// 入力に含まれる値は全て整数である

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

	var A, B, C, D int
	fmt.Fscan(in, &A, &B, &C, &D)

	sieves := sieve(200)

	idx := 0
	for n := A; n <= B; n++ {
		for {
			if sieves[idx] - n > D {
				fmt.Fprintln(out, "Takahashi")
				return
			}

			if sieves[idx] - n >= C && sieves[idx] - n <= D {
				break
			}

			idx++
		}
	}

	fmt.Fprintln(out, "Aoki")
}

// 素数リスト
func sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
