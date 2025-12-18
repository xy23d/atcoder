// https://atcoder.jp/contests/adt_hard_20251217_1/tasks/abc304_d

// 問題文
// xy -平面上にいくつかのイチゴが載った長方形のケーキがあります。 ケーキは、長方形領域 
// {(x,y):0≤x≤W,0≤y≤H} をちょうど占めます。

// ケーキには 
// N 個のイチゴが載っており、
// i=1,2,…,N について、
// i 番目のイチゴの座標は 
// (p 
// i
// ​
//  ,q 
// i
// ​
//  ) です。 
// 2 個以上のイチゴが同一の座標にあることはありません。

// 高橋君は、このケーキを包丁で下記の通りにいくつかのピースに切り分けます。

// まず、ケーキを通る 
// y 軸に並行な 
// A 本の異なる直線、直線 
// x=a 
// 1
// ​
//   、直線 
// x=a 
// 2
// ​
//   、
// … 、直線 
// x=a 
// A
// ​
//   のそれぞれにそってケーキを切る。
// 次に、ケーキを通る 
// x 軸に並行な 
// B 本の異なる直線、直線 
// y=b 
// 1
// ​
//   、直線 
// y=b 
// 2
// ​
//   、
// … 、直線 
// y=b 
// B
// ​
//   のそれぞれにそってケーキを切る。
// その結果、ケーキは 
// (A+1)(B+1) 個の長方形のピースに分割されます。 高橋君はそれらのうちのいずれか 
// 1 個だけを選んで食べます。 高橋君が選んだピースに載っているイチゴの個数としてあり得る最小値と最大値をそれぞれ出力してください。

// ここで、最終的にピースの縁となる位置にはイチゴが存在しないことが保証されます。 より形式的な説明は下記の制約を参照してください。

// 制約
// 3≤W,H≤10 
// 9
 
// 1≤N≤2×10 
// 5
 
// 0<p 
// i
// ​
//  <W
// 0<q 
// i
// ​
//  <H
// i
// 
// =j⟹(p 
// i
// ​
//  ,q 
// i
// ​
//  )
// 
// =(p 
// j
// ​
//  ,q 
// j
// ​
//  )
// 1≤A,B≤2×10 
// 5
 
// 0<a 
// 1
// ​
//  <a 
// 2
// ​
//  <⋯<a 
// A
// ​
//  <W
// 0<b 
// 1
// ​
//  <b 
// 2
// ​
//  <⋯<b 
// B
// ​
//  <H
// p 
// i
// ​
 
// 
// ∈{a 
// 1
// ​
//  ,a 
// 2
// ​
//  ,…,a 
// A
// ​
//  }
// q 
// i
// ​
 
// 
// ∈{b 
// 1
// ​
//  ,b 
// 2
// ​
//  ,…,b 
// B
// ​
//  }
// 入力はすべて整数

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var W, H, N int
	fmt.Fscan(in, &W, &H, &N)

	// FIXME: W,H は最大 1e9 なのでこの二次元配列は確保不能（メモリ爆発する）
	pq := make([][]bool, W+1)
	for i := 1; i <= W; i++ {
		pq[i] = make([]bool, H+1)
	}

	for i := 0; i < N; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		pq[p][q] = true
	}

	var A int
	fmt.Fscan(in, &A)
	as := make([]int, A)
	for i := 0; i < A; i++ {
		fmt.Fscan(in, &as[i])
	}

	var B int
	fmt.Fscan(in, &B)
	bs := make([]int, B)
	for i := 0; i < B; i++ {
		fmt.Fscan(in, &bs[i])
	}

	sort.Ints(as)
	sort.Ints(bs)

	min := N
	max := 0
	ba := 1
	bb := 1
	// FIXME: 末尾区間（a_A < x < W, b_B < y < H）が検査されず (A+1)(B+1) ピースを網羅できていない
	for i := 0; i < len(as); i++ {
		for j := 0; j < len(bs); j++ {
			t := 0
			// FIXME: グリッド総当たりは W×H ≒ 1e18 反復となり到底終わらない
			for a := ba; a <= as[i]; a++ {
				for b := bb; b <= bs[j]; b++ {
					if pq[a][b] {
						t++
					}
				}
			}

			if min > t {
				min = t
			}

			if max < t {
				max = t
			}

			bb = bs[j]
		}

		ba = as[i]
		bb = 1
	}

	fmt.Fprintln(out, min, max)
}
