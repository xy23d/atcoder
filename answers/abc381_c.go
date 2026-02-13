// 問題文
// この問題における 11/22 文字列の定義は A 問題および E 問題と同じです。

// 文字列 
// T が以下の条件を全て満たすとき、
// T を 11/22 文字列 と呼びます。

// ∣T∣ は奇数である。ここで、
// ∣T∣ は 
// T の長さを表す。
// 1 文字目から 
// 2
// ∣T∣+1
// ​
//  −1 文字目までが 1 である。
// 2
// ∣T∣+1
// ​
//   文字目が / である。
// 2
// ∣T∣+1
// ​
//  +1 文字目から 
// ∣T∣ 文字目までが 2 である。
// 例えば 11/22, 111/222, / は 11/22 文字列ですが、1122, 1/22, 11/2222, 22/11, //2/2/211 はそうではありません。

// 1, 2, / からなる長さ 
// N の文字列 
// S が与えられます。
// S は / を 
// 1 個以上含みます。
// 11/22 文字列であるような 
// S の(連続な)部分文字列の長さの最大値を求めてください。

// 制約
// 1≤N≤2×10 
// 5
 
// S は 1, 2, / からなる長さ 
// N の文字列
// S は / を 
// 1 個以上含む

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	var S string
	fmt.Fscan(in, &S)

	chars := strings.Split(S, "")

	ans := 0
	for i := 0; i < N; i++ {
		if chars[i] == "/" {
			t := execute(chars, i)
			if ans < t {
				ans = t
			}
		}
	}
	fmt.Fprintln(out, ans)
}

func execute(chars []string, idx int) int {
	loop := 0

	for {
		if idx-(loop+1) < 0 {
			break
		}

		if idx+(loop+1) >= len(chars) {
			break
		}

		if chars[idx-(loop+1)] != "1" {
			break
		}

		if chars[idx+(loop+1)] != "2" {
			break
		}

		loop++
	}

	if loop == 0 {
		return 1
	}

	return loop*2 + 1
}
