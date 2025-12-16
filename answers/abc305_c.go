// https://atcoder.jp/contests/adt_hard_20251216_1/tasks/abc305_c
// 
// 問題文
// 縦 
// H マス, 横 
// W マスのグリッドがあります。上から 
// i 行目、左から 
// j 列目のマスを 
// (i,j) と呼びます。
// はじめ、グリッド上には、ある 縦横 
// 2 マス以上 の部分長方形の内部にあるマスにクッキーが 1 枚ずつ置かれていて、それ以外のマスにはクッキーが置かれていません。
// 形式的に説明すると、以下の条件を全て満たす 4 つの整数の組 
// (a,b,c,d) がただ 1 つ存在します。

// 1≤a<b≤H
// 1≤c<d≤W
// グリッド上のマスのうち、
// a≤i≤b,c≤j≤d を満たす全てのマス 
// (i,j) にはクッキーが 1 枚ずつ置かれていて、それ以外のマスにはクッキーが置かれていない。
// ところが、すぬけ君がグリッド上のクッキーのどれか 1 枚を取って食べてしまいました。
// すぬけ君がクッキーを取ったマスは、クッキーが置かれていない状態に変わります。

// すぬけ君がクッキーを食べた後のグリッドの状態が入力として与えられます。
// マス 
// (i,j) の状態は文字 
// S 
// i,j
// ​
//   として与えられて、# はクッキーが置かれているマスを, . はクッキーが置かれていないマスを意味します。
// すぬけ君が食べたクッキーが元々置かれていたマスを答えてください。(答えは一意に定まります。)

// 制約
// 2≤H,W≤500
// S 
// i,j
// ​
//   は # または .

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

	var H, W int
	fmt.Fscan(in, &H, &W)

	isExists := false
	minW := H
	maxW := 0
	for i := 0; i < H; i++ {
		var S string
		fmt.Fscan(in, &S)
		chars := strings.Split(S, "")

		if !isExists {
			isBlankExists := false

			for j := 0; j < W; j++ {
				if chars[j] == "#" {
					isExists = true

					if minW > j {
						minW = j
					}

					if maxW < j {
						maxW = j
					}

					if isBlankExists {
						fmt.Fprintln(out, i+1, j)
						return
					}
				} else {
					if minW < j && isExists {
						isBlankExists = true
					}
				}
			}
		} else {
			for j := 0; j < W; j++ {
				if chars[j] == "#" {
					if minW > j {
						fmt.Fprintln(out, i, j+1)
						return
					}

					if maxW < j {
						fmt.Fprintln(out, i, j+1)
						return
					}
				} else {
					if minW <= j && j <= maxW {
						fmt.Fprintln(out, i+1, j+1)
						return
					}
				}
			}
		}
	}
}
